// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: wiseye_rpc.proto

package wiseye_rpc

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
	WiseyeGrpc_SayHello_FullMethodName   = "/wiseye_rpc.WiseyeGrpc/SayHello"
	WiseyeGrpc_UploadFile_FullMethodName = "/wiseye_rpc.WiseyeGrpc/UploadFile"
)

// WiseyeGrpcClient is the client API for WiseyeGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WiseyeGrpcClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	UploadFile(ctx context.Context, in *FileContent, opts ...grpc.CallOption) (*Empty, error)
}

type wiseyeGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewWiseyeGrpcClient(cc grpc.ClientConnInterface) WiseyeGrpcClient {
	return &wiseyeGrpcClient{cc}
}

func (c *wiseyeGrpcClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, WiseyeGrpc_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wiseyeGrpcClient) UploadFile(ctx context.Context, in *FileContent, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, WiseyeGrpc_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WiseyeGrpcServer is the server API for WiseyeGrpc service.
// All implementations must embed UnimplementedWiseyeGrpcServer
// for forward compatibility
type WiseyeGrpcServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	UploadFile(context.Context, *FileContent) (*Empty, error)
	mustEmbedUnimplementedWiseyeGrpcServer()
}

// UnimplementedWiseyeGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedWiseyeGrpcServer struct {
}

func (UnimplementedWiseyeGrpcServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedWiseyeGrpcServer) UploadFile(context.Context, *FileContent) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedWiseyeGrpcServer) mustEmbedUnimplementedWiseyeGrpcServer() {}

// UnsafeWiseyeGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WiseyeGrpcServer will
// result in compilation errors.
type UnsafeWiseyeGrpcServer interface {
	mustEmbedUnimplementedWiseyeGrpcServer()
}

func RegisterWiseyeGrpcServer(s grpc.ServiceRegistrar, srv WiseyeGrpcServer) {
	s.RegisterService(&WiseyeGrpc_ServiceDesc, srv)
}

func _WiseyeGrpc_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WiseyeGrpcServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WiseyeGrpc_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WiseyeGrpcServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WiseyeGrpc_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WiseyeGrpcServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WiseyeGrpc_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WiseyeGrpcServer).UploadFile(ctx, req.(*FileContent))
	}
	return interceptor(ctx, in, info, handler)
}

// WiseyeGrpc_ServiceDesc is the grpc.ServiceDesc for WiseyeGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WiseyeGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiseye_rpc.WiseyeGrpc",
	HandlerType: (*WiseyeGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _WiseyeGrpc_SayHello_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _WiseyeGrpc_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wiseye_rpc.proto",
}