// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// BizClient is the client API for Biz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BizClient interface {
	GetUsers(ctx context.Context, in *GetUserInput1, opts ...grpc.CallOption) (*GetUserOutput, error)
	GetUsersWithSqlInjection(ctx context.Context, in *GetUserInput2, opts ...grpc.CallOption) (*GetUserOutput, error)
}

type bizClient struct {
	cc grpc.ClientConnInterface
}

func NewBizClient(cc grpc.ClientConnInterface) BizClient {
	return &bizClient{cc}
}

func (c *bizClient) GetUsers(ctx context.Context, in *GetUserInput1, opts ...grpc.CallOption) (*GetUserOutput, error) {
	out := new(GetUserOutput)
	err := c.cc.Invoke(ctx, "/auth.Biz/get_users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizClient) GetUsersWithSqlInjection(ctx context.Context, in *GetUserInput2, opts ...grpc.CallOption) (*GetUserOutput, error) {
	out := new(GetUserOutput)
	err := c.cc.Invoke(ctx, "/auth.Biz/get_users_with_sql_injection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BizServer is the server API for Biz service.
// All implementations must embed UnimplementedBizServer
// for forward compatibility
type BizServer interface {
	GetUsers(context.Context, *GetUserInput1) (*GetUserOutput, error)
	GetUsersWithSqlInjection(context.Context, *GetUserInput2) (*GetUserOutput, error)
	mustEmbedUnimplementedBizServer()
}

// UnimplementedBizServer must be embedded to have forward compatible implementations.
type UnimplementedBizServer struct {
}

func (UnimplementedBizServer) GetUsers(context.Context, *GetUserInput1) (*GetUserOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedBizServer) GetUsersWithSqlInjection(context.Context, *GetUserInput2) (*GetUserOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersWithSqlInjection not implemented")
}
func (UnimplementedBizServer) mustEmbedUnimplementedBizServer() {}

// UnsafeBizServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BizServer will
// result in compilation errors.
type UnsafeBizServer interface {
	mustEmbedUnimplementedBizServer()
}

func RegisterBizServer(s grpc.ServiceRegistrar, srv BizServer) {
	s.RegisterService(&Biz_ServiceDesc, srv)
}

func _Biz_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInput1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Biz/get_users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServer).GetUsers(ctx, req.(*GetUserInput1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Biz_GetUsersWithSqlInjection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInput2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServer).GetUsersWithSqlInjection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Biz/get_users_with_sql_injection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServer).GetUsersWithSqlInjection(ctx, req.(*GetUserInput2))
	}
	return interceptor(ctx, in, info, handler)
}

// Biz_ServiceDesc is the grpc.ServiceDesc for Biz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Biz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Biz",
	HandlerType: (*BizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_users",
			Handler:    _Biz_GetUsers_Handler,
		},
		{
			MethodName: "get_users_with_sql_injection",
			Handler:    _Biz_GetUsersWithSqlInjection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
