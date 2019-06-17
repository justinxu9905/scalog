package data

import (
	"context"
	"io"

	"github.com/scalog/scalogger/data/datapb"
	log "github.com/scalog/scalogger/logger"
)

func (server *DataServer) Append(stream datapb.Data_AppendServer) error {
	initialized := false
	done := make(chan struct{})
	for {
		select {
		case <-done:
			return nil
		default:
			record, err := stream.Recv()
			if err != nil {
				close(done)
				if err == io.EOF {
					return nil
				}
				return err
			}
			if !initialized {
				cid := record.ClientID
				go server.respondToClient(cid, done, stream)
				initialized = true
			}
			server.appendC <- record

		}
	}
}

func (server *DataServer) respondToClient(cid int32, done chan struct{}, stream datapb.Data_AppendServer) {
	ackC := make(chan *datapb.Ack)
	server.ackCMu.Lock()
	server.ackC[cid] = ackC
	server.ackCMu.Unlock()
	for {
		select {
		case <-done:
			server.ackCMu.Lock()
			delete(server.ackC, cid)
			server.ackCMu.Unlock()
			log.Infof("Client %v is closed", cid)
			close(ackC)
			return
		case ack := <-ackC:
			if err := stream.Send(ack); err != nil {
				server.ackCMu.Lock()
				delete(server.ackC, cid)
				server.ackCMu.Unlock()
				log.Infof("Client %v is closed", cid)
				close(ackC)
				close(done)
				return
			}
		}
	}
}

func (server *DataServer) Replicate(stream datapb.Data_AppendServer) error {
	for {
		select {
		default:
			record, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			server.replicateC <- record
		}
	}
}

// TODO implement the trim operation
func (server *DataServer) Trim(ctx context.Context, gsn *datapb.GlobalSN) (*datapb.Empty, error) {
	return &datapb.Empty{}, nil
}

func (server *DataServer) Read(ctx context.Context, gsn *datapb.GlobalSN) (*datapb.Record, error) {
	r, rid, err := server.storage.Read(gsn.Gsn)
	if err != nil {
		return nil, err
	}
	record := &datapb.Record{
		GlobalSN:       gsn.Gsn,
		ShardID:        server.shardID,
		LocalReplicaID: rid,
		ViewID:         server.viewID,
		Record:         r,
	}
	return record, err
}

func (server *DataServer) Subscribe(gsn *datapb.GlobalSN, stream datapb.Data_SubscribeServer) error {
	subC := make(chan *datapb.Record)
	server.subCMu.Lock()
	cid := server.clientID
	server.clientID++
	server.subC[cid] = subC
	server.subCMu.Unlock()
	for {
		select {
		case sub := <-subC:
			if err := stream.Send(sub); err != nil {
				server.subCMu.Lock()
				delete(server.subC, cid)
				server.subCMu.Unlock()
				log.Infof("Client %v is closed", cid)
				close(subC)
				return err
			}
		}
	}
	return nil
}