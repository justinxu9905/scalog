// Code generated by protoc-gen-go.
// source: orderpb/order.proto
// DO NOT EDIT!

/*
Package orderpb is a generated protocol buffer package.

It is generated from these files:
	orderpb/order.proto

It has these top-level messages:
	LocalCut
	LocalCuts
	CommittedCut
	CommittedEntry
	FinalizeEntry
	Empty
*/
package orderpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocalCut struct {
	ShardID        int32   `protobuf:"varint,1,opt,name=shardID" json:"shardID,omitempty"`
	LocalReplicaID int32   `protobuf:"varint,2,opt,name=localReplicaID" json:"localReplicaID,omitempty"`
	Cut            []int32 `protobuf:"varint,3,rep,packed,name=cut" json:"cut,omitempty"`
}

func (m *LocalCut) Reset()                    { *m = LocalCut{} }
func (m *LocalCut) String() string            { return proto.CompactTextString(m) }
func (*LocalCut) ProtoMessage()               {}
func (*LocalCut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocalCut) GetShardID() int32 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *LocalCut) GetLocalReplicaID() int32 {
	if m != nil {
		return m.LocalReplicaID
	}
	return 0
}

func (m *LocalCut) GetCut() []int32 {
	if m != nil {
		return m.Cut
	}
	return nil
}

type LocalCuts struct {
	Cuts []*LocalCut `protobuf:"bytes,1,rep,name=cuts" json:"cuts,omitempty"`
}

func (m *LocalCuts) Reset()                    { *m = LocalCuts{} }
func (m *LocalCuts) String() string            { return proto.CompactTextString(m) }
func (*LocalCuts) ProtoMessage()               {}
func (*LocalCuts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LocalCuts) GetCuts() []*LocalCut {
	if m != nil {
		return m.Cuts
	}
	return nil
}

type CommittedCut struct {
	StartGSN int32 `protobuf:"varint,1,opt,name=startGSN" json:"startGSN,omitempty"`
	// from globalReplicaID to each cut entry
	// globalReplicaID = shardID * numReplicas + localReplicaID
	Cut map[int32]int32 `protobuf:"bytes,2,rep,name=cut" json:"cut,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *CommittedCut) Reset()                    { *m = CommittedCut{} }
func (m *CommittedCut) String() string            { return proto.CompactTextString(m) }
func (*CommittedCut) ProtoMessage()               {}
func (*CommittedCut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CommittedCut) GetStartGSN() int32 {
	if m != nil {
		return m.StartGSN
	}
	return 0
}

func (m *CommittedCut) GetCut() map[int32]int32 {
	if m != nil {
		return m.Cut
	}
	return nil
}

type CommittedEntry struct {
	Seq            int64          `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	ViewID         int32          `protobuf:"varint,2,opt,name=viewID" json:"viewID,omitempty"`
	CommittedCut   *CommittedCut  `protobuf:"bytes,3,opt,name=committedCut" json:"committedCut,omitempty"`
	FinalizeShards *FinalizeEntry `protobuf:"bytes,4,opt,name=finalizeShards" json:"finalizeShards,omitempty"`
}

func (m *CommittedEntry) Reset()                    { *m = CommittedEntry{} }
func (m *CommittedEntry) String() string            { return proto.CompactTextString(m) }
func (*CommittedEntry) ProtoMessage()               {}
func (*CommittedEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CommittedEntry) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *CommittedEntry) GetViewID() int32 {
	if m != nil {
		return m.ViewID
	}
	return 0
}

func (m *CommittedEntry) GetCommittedCut() *CommittedCut {
	if m != nil {
		return m.CommittedCut
	}
	return nil
}

func (m *CommittedEntry) GetFinalizeShards() *FinalizeEntry {
	if m != nil {
		return m.FinalizeShards
	}
	return nil
}

type FinalizeEntry struct {
	Limit    int32   `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	ShardIDs []int32 `protobuf:"varint,2,rep,packed,name=shardIDs" json:"shardIDs,omitempty"`
}

func (m *FinalizeEntry) Reset()                    { *m = FinalizeEntry{} }
func (m *FinalizeEntry) String() string            { return proto.CompactTextString(m) }
func (*FinalizeEntry) ProtoMessage()               {}
func (*FinalizeEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FinalizeEntry) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FinalizeEntry) GetShardIDs() []int32 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*LocalCut)(nil), "orderpb.LocalCut")
	proto.RegisterType((*LocalCuts)(nil), "orderpb.LocalCuts")
	proto.RegisterType((*CommittedCut)(nil), "orderpb.CommittedCut")
	proto.RegisterType((*CommittedEntry)(nil), "orderpb.CommittedEntry")
	proto.RegisterType((*FinalizeEntry)(nil), "orderpb.FinalizeEntry")
	proto.RegisterType((*Empty)(nil), "orderpb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Order service

type OrderClient interface {
	Report(ctx context.Context, opts ...grpc.CallOption) (Order_ReportClient, error)
	Forward(ctx context.Context, opts ...grpc.CallOption) (Order_ForwardClient, error)
	Finalize(ctx context.Context, in *FinalizeEntry, opts ...grpc.CallOption) (*Empty, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Report(ctx context.Context, opts ...grpc.CallOption) (Order_ReportClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Order_serviceDesc.Streams[0], c.cc, "/orderpb.Order/Report", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderReportClient{stream}
	return x, nil
}

type Order_ReportClient interface {
	Send(*LocalCuts) error
	Recv() (*CommittedEntry, error)
	grpc.ClientStream
}

type orderReportClient struct {
	grpc.ClientStream
}

func (x *orderReportClient) Send(m *LocalCuts) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderReportClient) Recv() (*CommittedEntry, error) {
	m := new(CommittedEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderClient) Forward(ctx context.Context, opts ...grpc.CallOption) (Order_ForwardClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Order_serviceDesc.Streams[1], c.cc, "/orderpb.Order/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderForwardClient{stream}
	return x, nil
}

type Order_ForwardClient interface {
	Send(*LocalCuts) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type orderForwardClient struct {
	grpc.ClientStream
}

func (x *orderForwardClient) Send(m *LocalCuts) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderForwardClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderClient) Finalize(ctx context.Context, in *FinalizeEntry, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/orderpb.Order/Finalize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderServer interface {
	Report(Order_ReportServer) error
	Forward(Order_ForwardServer) error
	Finalize(context.Context, *FinalizeEntry) (*Empty, error)
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_Report_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServer).Report(&orderReportServer{stream})
}

type Order_ReportServer interface {
	Send(*CommittedEntry) error
	Recv() (*LocalCuts, error)
	grpc.ServerStream
}

type orderReportServer struct {
	grpc.ServerStream
}

func (x *orderReportServer) Send(m *CommittedEntry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderReportServer) Recv() (*LocalCuts, error) {
	m := new(LocalCuts)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Order_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServer).Forward(&orderForwardServer{stream})
}

type Order_ForwardServer interface {
	SendAndClose(*Empty) error
	Recv() (*LocalCuts, error)
	grpc.ServerStream
}

type orderForwardServer struct {
	grpc.ServerStream
}

func (x *orderForwardServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderForwardServer) Recv() (*LocalCuts, error) {
	m := new(LocalCuts)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Order_Finalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Finalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderpb.Order/Finalize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Finalize(ctx, req.(*FinalizeEntry))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderpb.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Finalize",
			Handler:    _Order_Finalize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Report",
			Handler:       _Order_Report_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Forward",
			Handler:       _Order_Forward_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "orderpb/order.proto",
}

func init() { proto.RegisterFile("orderpb/order.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xdd, 0xca, 0xd3, 0x40,
	0x10, 0x75, 0x9b, 0x2f, 0x4d, 0x9c, 0xef, 0x33, 0xe8, 0xa8, 0x35, 0xf4, 0x42, 0x42, 0x40, 0xc9,
	0x55, 0xac, 0x51, 0xc4, 0x1f, 0x10, 0xa4, 0x3f, 0x52, 0x10, 0x85, 0xf4, 0x5e, 0x48, 0x93, 0x15,
	0x83, 0x49, 0x13, 0x77, 0x37, 0x2d, 0xf5, 0x39, 0x7c, 0x0d, 0x2f, 0x7d, 0x3f, 0xd9, 0xcd, 0x26,
	0xf4, 0xcf, 0xab, 0xec, 0x99, 0x3d, 0x67, 0xe6, 0x9c, 0xd9, 0xc0, 0xfd, 0x8a, 0x65, 0x94, 0xd5,
	0xeb, 0x67, 0xea, 0x1b, 0xd6, 0xac, 0x12, 0x15, 0x5a, 0xba, 0xe8, 0x7f, 0x05, 0xfb, 0x53, 0x95,
	0x26, 0xc5, 0xb4, 0x11, 0xe8, 0x82, 0xc5, 0xbf, 0x27, 0x2c, 0x5b, 0xce, 0x5c, 0xe2, 0x91, 0xc0,
	0x8c, 0x3b, 0x88, 0x4f, 0xc1, 0x29, 0x24, 0x2b, 0xa6, 0x75, 0x91, 0xa7, 0xc9, 0x72, 0xe6, 0x0e,
	0x14, 0xe1, 0xa4, 0x8a, 0x77, 0xc1, 0x48, 0x1b, 0xe1, 0x1a, 0x9e, 0x11, 0x98, 0xb1, 0x3c, 0xfa,
	0x11, 0xdc, 0xee, 0xfa, 0x73, 0x7c, 0x02, 0x57, 0x69, 0x23, 0xb8, 0x4b, 0x3c, 0x23, 0xb8, 0x8e,
	0xee, 0x85, 0xda, 0x44, 0xd8, 0x31, 0x62, 0x75, 0xed, 0xff, 0x26, 0x70, 0x33, 0xad, 0xca, 0x32,
	0x17, 0x82, 0x66, 0xd2, 0xd8, 0x18, 0x6c, 0x2e, 0x12, 0x26, 0x3e, 0xae, 0x3e, 0x6b, 0x67, 0x3d,
	0xc6, 0x49, 0x3b, 0x72, 0xa0, 0x5a, 0x3e, 0xee, 0x5b, 0x1e, 0xea, 0xc3, 0x69, 0x23, 0xe6, 0x1b,
	0xc1, 0xf6, 0xca, 0xd2, 0xf8, 0x15, 0xd8, 0x5d, 0x41, 0x1a, 0xfe, 0x41, 0xf7, 0xba, 0xa9, 0x3c,
	0xe2, 0x03, 0x30, 0xb7, 0x49, 0xd1, 0x50, 0x9d, 0xb0, 0x05, 0x6f, 0x07, 0xaf, 0x89, 0xff, 0x97,
	0x80, 0xd3, 0xb7, 0xed, 0xe5, 0x9c, 0xfe, 0x54, 0x72, 0x23, 0x96, 0x47, 0x1c, 0xc1, 0x70, 0x9b,
	0xd3, 0x5d, 0xbf, 0x21, 0x8d, 0xf0, 0x0d, 0xdc, 0xa4, 0x07, 0x96, 0x5c, 0xc3, 0x23, 0xc1, 0x75,
	0xf4, 0xf0, 0xa2, 0xdf, 0xf8, 0x88, 0x8a, 0xef, 0xc1, 0xf9, 0x96, 0x6f, 0x92, 0x22, 0xff, 0x45,
	0x57, 0xf2, 0x3d, 0xb8, 0x7b, 0xa5, 0xc4, 0xa3, 0x5e, 0xbc, 0xd0, 0xd7, 0x6d, 0xc8, 0x13, 0xb6,
	0xff, 0x01, 0xee, 0x1c, 0x11, 0x64, 0xc4, 0x22, 0x2f, 0x73, 0xa1, 0x63, 0xb7, 0x40, 0x2d, 0xb9,
	0x7d, 0x6e, 0xae, 0xb6, 0x29, 0x97, 0xac, 0xb1, 0x6f, 0x81, 0x39, 0x2f, 0x6b, 0xb1, 0x8f, 0xfe,
	0x10, 0x30, 0xbf, 0xc8, 0xa9, 0xf8, 0x0e, 0x86, 0x31, 0xad, 0x2b, 0x26, 0x10, 0xcf, 0xde, 0x91,
	0x8f, 0x1f, 0x9d, 0x07, 0x53, 0xb3, 0xfd, 0x5b, 0x01, 0x99, 0x10, 0x7c, 0x0e, 0xd6, 0xa2, 0x62,
	0xbb, 0x84, 0x65, 0x17, 0xd5, 0x4e, 0x5f, 0x53, 0x53, 0xa5, 0x08, 0x5f, 0x82, 0xdd, 0xa5, 0xc0,
	0xff, 0x24, 0x3f, 0xd7, 0xad, 0x87, 0xea, 0x77, 0x7f, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xff, 0x8e, 0x26, 0x05, 0x03, 0x00, 0x00,
}