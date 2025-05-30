// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package comm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MultiPaxosRPCClient is the client API for MultiPaxosRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiPaxosRPCClient interface {
	Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*AcceptResponse, error)
	Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*PrepareResponse, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	ResumeSnapshot(ctx context.Context, opts ...grpc.CallOption) (MultiPaxosRPC_ResumeSnapshotClient, error)
	InstancesGap(ctx context.Context, in *InstanceRequest, opts ...grpc.CallOption) (MultiPaxosRPC_InstancesGapClient, error)
}

type multiPaxosRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiPaxosRPCClient(cc grpc.ClientConnInterface) MultiPaxosRPCClient {
	return &multiPaxosRPCClient{cc}
}

func (c *multiPaxosRPCClient) Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*AcceptResponse, error) {
	out := new(AcceptResponse)
	err := c.cc.Invoke(ctx, "/multipaxosrpc.MultiPaxosRPC/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiPaxosRPCClient) Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*PrepareResponse, error) {
	out := new(PrepareResponse)
	err := c.cc.Invoke(ctx, "/multipaxosrpc.MultiPaxosRPC/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiPaxosRPCClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/multipaxosrpc.MultiPaxosRPC/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiPaxosRPCClient) ResumeSnapshot(ctx context.Context, opts ...grpc.CallOption) (MultiPaxosRPC_ResumeSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MultiPaxosRPC_serviceDesc.Streams[0], "/multipaxosrpc.MultiPaxosRPC/ResumeSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiPaxosRPCResumeSnapshotClient{stream}
	return x, nil
}

type MultiPaxosRPC_ResumeSnapshotClient interface {
	Send(*SnapshotRequest) error
	CloseAndRecv() (*SnapshotResponse, error)
	grpc.ClientStream
}

type multiPaxosRPCResumeSnapshotClient struct {
	grpc.ClientStream
}

func (x *multiPaxosRPCResumeSnapshotClient) Send(m *SnapshotRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiPaxosRPCResumeSnapshotClient) CloseAndRecv() (*SnapshotResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *multiPaxosRPCClient) InstancesGap(ctx context.Context, in *InstanceRequest, opts ...grpc.CallOption) (MultiPaxosRPC_InstancesGapClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MultiPaxosRPC_serviceDesc.Streams[1], "/multipaxosrpc.MultiPaxosRPC/InstancesGap", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiPaxosRPCInstancesGapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MultiPaxosRPC_InstancesGapClient interface {
	Recv() (*Instance, error)
	grpc.ClientStream
}

type multiPaxosRPCInstancesGapClient struct {
	grpc.ClientStream
}

func (x *multiPaxosRPCInstancesGapClient) Recv() (*Instance, error) {
	m := new(Instance)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiPaxosRPCServer is the server API for MultiPaxosRPC service.
// All implementations must embed UnimplementedMultiPaxosRPCServer
// for forward compatibility
type MultiPaxosRPCServer interface {
	Accept(context.Context, *AcceptRequest) (*AcceptResponse, error)
	Prepare(context.Context, *PrepareRequest) (*PrepareResponse, error)
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	ResumeSnapshot(MultiPaxosRPC_ResumeSnapshotServer) error
	InstancesGap(*InstanceRequest, MultiPaxosRPC_InstancesGapServer) error
	mustEmbedUnimplementedMultiPaxosRPCServer()
}

// UnimplementedMultiPaxosRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMultiPaxosRPCServer struct {
}

func (*UnimplementedMultiPaxosRPCServer) Accept(context.Context, *AcceptRequest) (*AcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (*UnimplementedMultiPaxosRPCServer) Prepare(context.Context, *PrepareRequest) (*PrepareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedMultiPaxosRPCServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedMultiPaxosRPCServer) ResumeSnapshot(MultiPaxosRPC_ResumeSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method ResumeSnapshot not implemented")
}
func (*UnimplementedMultiPaxosRPCServer) InstancesGap(*InstanceRequest, MultiPaxosRPC_InstancesGapServer) error {
	return status.Errorf(codes.Unimplemented, "method InstancesGap not implemented")
}
func (*UnimplementedMultiPaxosRPCServer) mustEmbedUnimplementedMultiPaxosRPCServer() {}

func RegisterMultiPaxosRPCServer(s *grpc.Server, srv MultiPaxosRPCServer) {
	s.RegisterService(&_MultiPaxosRPC_serviceDesc, srv)
}

func _MultiPaxosRPC_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiPaxosRPCServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multipaxosrpc.MultiPaxosRPC/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiPaxosRPCServer).Accept(ctx, req.(*AcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiPaxosRPC_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiPaxosRPCServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multipaxosrpc.MultiPaxosRPC/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiPaxosRPCServer).Prepare(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiPaxosRPC_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiPaxosRPCServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multipaxosrpc.MultiPaxosRPC/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiPaxosRPCServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultiPaxosRPC_ResumeSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiPaxosRPCServer).ResumeSnapshot(&multiPaxosRPCResumeSnapshotServer{stream})
}

type MultiPaxosRPC_ResumeSnapshotServer interface {
	SendAndClose(*SnapshotResponse) error
	Recv() (*SnapshotRequest, error)
	grpc.ServerStream
}

type multiPaxosRPCResumeSnapshotServer struct {
	grpc.ServerStream
}

func (x *multiPaxosRPCResumeSnapshotServer) SendAndClose(m *SnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiPaxosRPCResumeSnapshotServer) Recv() (*SnapshotRequest, error) {
	m := new(SnapshotRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MultiPaxosRPC_InstancesGap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstanceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MultiPaxosRPCServer).InstancesGap(m, &multiPaxosRPCInstancesGapServer{stream})
}

type MultiPaxosRPC_InstancesGapServer interface {
	Send(*Instance) error
	grpc.ServerStream
}

type multiPaxosRPCInstancesGapServer struct {
	grpc.ServerStream
}

func (x *multiPaxosRPCInstancesGapServer) Send(m *Instance) error {
	return x.ServerStream.SendMsg(m)
}

var _MultiPaxosRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multipaxosrpc.MultiPaxosRPC",
	HandlerType: (*MultiPaxosRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Accept",
			Handler:    _MultiPaxosRPC_Accept_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _MultiPaxosRPC_Prepare_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _MultiPaxosRPC_Commit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ResumeSnapshot",
			Handler:       _MultiPaxosRPC_ResumeSnapshot_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "InstancesGap",
			Handler:       _MultiPaxosRPC_InstancesGap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "multipaxosrpc.proto",
}
