// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mq_pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SeaweedMessagingClient is the client API for SeaweedMessaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeaweedMessagingClient interface {
	// control plane
	FindBrokerLeader(ctx context.Context, in *FindBrokerLeaderRequest, opts ...grpc.CallOption) (*FindBrokerLeaderResponse, error)
	AssignSegmentBrokers(ctx context.Context, in *AssignSegmentBrokersRequest, opts ...grpc.CallOption) (*AssignSegmentBrokersResponse, error)
	CheckSegmentStatus(ctx context.Context, in *CheckSegmentStatusRequest, opts ...grpc.CallOption) (*CheckSegmentStatusResponse, error)
	CheckBrokerLoad(ctx context.Context, in *CheckBrokerLoadRequest, opts ...grpc.CallOption) (*CheckBrokerLoadResponse, error)
	// data plane
	Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error)
}

type seaweedMessagingClient struct {
	cc grpc.ClientConnInterface
}

func NewSeaweedMessagingClient(cc grpc.ClientConnInterface) SeaweedMessagingClient {
	return &seaweedMessagingClient{cc}
}

func (c *seaweedMessagingClient) FindBrokerLeader(ctx context.Context, in *FindBrokerLeaderRequest, opts ...grpc.CallOption) (*FindBrokerLeaderResponse, error) {
	out := new(FindBrokerLeaderResponse)
	err := c.cc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/FindBrokerLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) AssignSegmentBrokers(ctx context.Context, in *AssignSegmentBrokersRequest, opts ...grpc.CallOption) (*AssignSegmentBrokersResponse, error) {
	out := new(AssignSegmentBrokersResponse)
	err := c.cc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/AssignSegmentBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) CheckSegmentStatus(ctx context.Context, in *CheckSegmentStatusRequest, opts ...grpc.CallOption) (*CheckSegmentStatusResponse, error) {
	out := new(CheckSegmentStatusResponse)
	err := c.cc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/CheckSegmentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) CheckBrokerLoad(ctx context.Context, in *CheckBrokerLoadRequest, opts ...grpc.CallOption) (*CheckBrokerLoadResponse, error) {
	out := new(CheckBrokerLoadResponse)
	err := c.cc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/CheckBrokerLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedMessaging_ServiceDesc.Streams[0], "/messaging_pb.SeaweedMessaging/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingPublishClient{stream}
	return x, nil
}

type SeaweedMessaging_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type seaweedMessagingPublishClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingPublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingPublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SeaweedMessagingServer is the server API for SeaweedMessaging service.
// All implementations must embed UnimplementedSeaweedMessagingServer
// for forward compatibility
type SeaweedMessagingServer interface {
	// control plane
	FindBrokerLeader(context.Context, *FindBrokerLeaderRequest) (*FindBrokerLeaderResponse, error)
	AssignSegmentBrokers(context.Context, *AssignSegmentBrokersRequest) (*AssignSegmentBrokersResponse, error)
	CheckSegmentStatus(context.Context, *CheckSegmentStatusRequest) (*CheckSegmentStatusResponse, error)
	CheckBrokerLoad(context.Context, *CheckBrokerLoadRequest) (*CheckBrokerLoadResponse, error)
	// data plane
	Publish(SeaweedMessaging_PublishServer) error
	mustEmbedUnimplementedSeaweedMessagingServer()
}

// UnimplementedSeaweedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedSeaweedMessagingServer struct {
}

func (UnimplementedSeaweedMessagingServer) FindBrokerLeader(context.Context, *FindBrokerLeaderRequest) (*FindBrokerLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBrokerLeader not implemented")
}
func (UnimplementedSeaweedMessagingServer) AssignSegmentBrokers(context.Context, *AssignSegmentBrokersRequest) (*AssignSegmentBrokersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignSegmentBrokers not implemented")
}
func (UnimplementedSeaweedMessagingServer) CheckSegmentStatus(context.Context, *CheckSegmentStatusRequest) (*CheckSegmentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSegmentStatus not implemented")
}
func (UnimplementedSeaweedMessagingServer) CheckBrokerLoad(context.Context, *CheckBrokerLoadRequest) (*CheckBrokerLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBrokerLoad not implemented")
}
func (UnimplementedSeaweedMessagingServer) Publish(SeaweedMessaging_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedSeaweedMessagingServer) mustEmbedUnimplementedSeaweedMessagingServer() {}

// UnsafeSeaweedMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeaweedMessagingServer will
// result in compilation errors.
type UnsafeSeaweedMessagingServer interface {
	mustEmbedUnimplementedSeaweedMessagingServer()
}

func RegisterSeaweedMessagingServer(s grpc.ServiceRegistrar, srv SeaweedMessagingServer) {
	s.RegisterService(&SeaweedMessaging_ServiceDesc, srv)
}

func _SeaweedMessaging_FindBrokerLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBrokerLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).FindBrokerLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/FindBrokerLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).FindBrokerLeader(ctx, req.(*FindBrokerLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_AssignSegmentBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignSegmentBrokersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).AssignSegmentBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/AssignSegmentBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).AssignSegmentBrokers(ctx, req.(*AssignSegmentBrokersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_CheckSegmentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSegmentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).CheckSegmentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/CheckSegmentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).CheckSegmentStatus(ctx, req.(*CheckSegmentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_CheckBrokerLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBrokerLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).CheckBrokerLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/CheckBrokerLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).CheckBrokerLoad(ctx, req.(*CheckBrokerLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Publish(&seaweedMessagingPublishServer{stream})
}

type SeaweedMessaging_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type seaweedMessagingPublishServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingPublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingPublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SeaweedMessaging_ServiceDesc is the grpc.ServiceDesc for SeaweedMessaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeaweedMessaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messaging_pb.SeaweedMessaging",
	HandlerType: (*SeaweedMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBrokerLeader",
			Handler:    _SeaweedMessaging_FindBrokerLeader_Handler,
		},
		{
			MethodName: "AssignSegmentBrokers",
			Handler:    _SeaweedMessaging_AssignSegmentBrokers_Handler,
		},
		{
			MethodName: "CheckSegmentStatus",
			Handler:    _SeaweedMessaging_CheckSegmentStatus_Handler,
		},
		{
			MethodName: "CheckBrokerLoad",
			Handler:    _SeaweedMessaging_CheckBrokerLoad_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _SeaweedMessaging_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mq.proto",
}