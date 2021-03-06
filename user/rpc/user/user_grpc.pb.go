// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	GetUserByPhone(ctx context.Context, in *GetByPhoneReq, opts ...grpc.CallOption) (*GetUserResp, error)
	GetUserByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*GetUserResp, error)
	VerifyPassword(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/add.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/add.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByPhone(ctx context.Context, in *GetByPhoneReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/add.User/getUserByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/add.User/getUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyPassword(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyResp, error) {
	out := new(VerifyResp)
	err := c.cc.Invoke(ctx, "/add.User/verifyPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	GetUserByPhone(context.Context, *GetByPhoneReq) (*GetUserResp, error)
	GetUserByEmail(context.Context, *GetByEmailReq) (*GetUserResp, error)
	VerifyPassword(context.Context, *VerifyReq) (*VerifyResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) GetUserByPhone(context.Context, *GetByPhoneReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByPhone not implemented")
}
func (UnimplementedUserServer) GetUserByEmail(context.Context, *GetByEmailReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserServer) VerifyPassword(context.Context, *VerifyReq) (*VerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPassword not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.User/getUserByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByPhone(ctx, req.(*GetByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.User/getUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByEmail(ctx, req.(*GetByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.User/verifyPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyPassword(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "add.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "getUserByPhone",
			Handler:    _User_GetUserByPhone_Handler,
		},
		{
			MethodName: "getUserByEmail",
			Handler:    _User_GetUserByEmail_Handler,
		},
		{
			MethodName: "verifyPassword",
			Handler:    _User_VerifyPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
