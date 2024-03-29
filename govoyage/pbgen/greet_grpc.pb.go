// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: greet.proto

package pbgen

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

const (
	Greeter_Greet_FullMethodName       = "/proto.Greeter/Greet"
	Greeter_GreetStream_FullMethodName = "/proto.Greeter/GreetStream"
	Greeter_GreetRecord_FullMethodName = "/proto.Greeter/GreetRecord"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (Greeter_GreetStreamClient, error)
	GreetRecord(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetRecordClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, Greeter_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GreetStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (Greeter_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], Greeter_GreetStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGreetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GreetStreamClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greeterGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greeterGreetStreamClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) GreetRecord(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], Greeter_GreetRecord_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGreetRecordClient{stream}
	return x, nil
}

type Greeter_GreetRecordClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greeterGreetRecordClient struct {
	grpc.ClientStream
}

func (x *greeterGreetRecordClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterGreetRecordClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetStream(*GreetRequest, Greeter_GreetStreamServer) error
	GreetRecord(Greeter_GreetRecordServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreeterServer) GreetStream(*GreetRequest, Greeter_GreetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetStream not implemented")
}
func (UnimplementedGreeterServer) GreetRecord(Greeter_GreetRecordServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetRecord not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GreetStream(m, &greeterGreetStreamServer{stream})
}

type Greeter_GreetStreamServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greeterGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greeterGreetStreamServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_GreetRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).GreetRecord(&greeterGreetRecordServer{stream})
}

type Greeter_GreetRecordServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greeterGreetRecordServer struct {
	grpc.ServerStream
}

func (x *greeterGreetRecordServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterGreetRecordServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Greeter_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetStream",
			Handler:       _Greeter_GreetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetRecord",
			Handler:       _Greeter_GreetRecord_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
