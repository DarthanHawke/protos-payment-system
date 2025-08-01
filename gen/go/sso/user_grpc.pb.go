// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0
// source: sso/user.proto

package pmtstmv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_Register_FullMethodName       = "/sso.UserService/Register"
	UserService_Login_FullMethodName          = "/sso.UserService/Login"
	UserService_GetProfile_FullMethodName     = "/sso.UserService/GetProfile"
	UserService_GetAllUsers_FullMethodName    = "/sso.UserService/GetAllUsers"
	UserService_UpdateName_FullMethodName     = "/sso.UserService/UpdateName"
	UserService_UpdateEmail_FullMethodName    = "/sso.UserService/UpdateEmail"
	UserService_UpdatePassword_FullMethodName = "/sso.UserService/UpdatePassword"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, UserService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	UpdateName(context.Context, *UpdateNameRequest) (*emptypb.Empty, error)
	UpdateEmail(context.Context, *UpdateEmailRequest) (*emptypb.Empty, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) UpdateName(context.Context, *UpdateNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedUserServiceServer) UpdateEmail(context.Context, *UpdateEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedUserServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateName(ctx, req.(*UpdateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateEmail(ctx, req.(*UpdateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _UserService_UpdateName_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _UserService_UpdateEmail_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/user.proto",
}
