// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: secrets.proto

package msi

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
	SecretService_GetSecret_FullMethodName = "/proto.SecretService/GetSecret"
)

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, SecretService_GetSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretService_GetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretService_GetSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secrets.proto",
}
