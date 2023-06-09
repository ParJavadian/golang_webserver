// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth.proto

package main

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	ReqPq(ctx context.Context, in *ReqPqInput, opts ...grpc.CallOption) (*ReqPqResponse, error)
	Req_DHParams(ctx context.Context, in *Req_DHParamsInput, opts ...grpc.CallOption) (*Req_DHParamsResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) ReqPq(ctx context.Context, in *ReqPqInput, opts ...grpc.CallOption) (*ReqPqResponse, error) {
	out := new(ReqPqResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/req_pq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Req_DHParams(ctx context.Context, in *Req_DHParamsInput, opts ...grpc.CallOption) (*Req_DHParamsResponse, error) {
	out := new(Req_DHParamsResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/req_DH_params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	ReqPq(context.Context, *ReqPqInput) (*ReqPqResponse, error)
	Req_DHParams(context.Context, *Req_DHParamsInput) (*Req_DHParamsResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) ReqPq(context.Context, *ReqPqInput) (*ReqPqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqPq not implemented")
}
func (UnimplementedAuthServer) Req_DHParams(context.Context, *Req_DHParamsInput) (*Req_DHParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Req_DHParams not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPqInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/req_pq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ReqPq(ctx, req.(*ReqPqInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Req_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req_DHParamsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Req_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/req_DH_params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Req_DHParams(ctx, req.(*Req_DHParamsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _Auth_ReqPq_Handler,
		},
		{
			MethodName: "req_DH_params",
			Handler:    _Auth_Req_DHParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
