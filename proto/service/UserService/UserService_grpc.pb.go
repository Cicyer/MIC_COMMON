// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package UserService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	PassLogin(ctx context.Context, in *UserPassLoginReq, opts ...grpc.CallOption) (*UserPassLoginResp, error)
	OauthToken(ctx context.Context, in *OauthTokenReq, opts ...grpc.CallOption) (*OauthTokenResp, error)
	GetOauthToken(ctx context.Context, in *GetOauthTokenReq, opts ...grpc.CallOption) (*OauthTokenResp, error)
	CreateOauthAccount(ctx context.Context, in *CreateOauthAccountReq, opts ...grpc.CallOption) (*CreateOauthAccountResp, error)
	//重置oauth 的secret
	ResetOauthAccountSecret(ctx context.Context, in *ResetOauthAccountReq, opts ...grpc.CallOption) (*CreateOauthAccountResp, error)
	//获取某个指定平台组织的oauth账号信息
	GetOauthAccount(ctx context.Context, in *GetOauthAccountReq, opts ...grpc.CallOption) (*GetOauthAccountResp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserResp, error)
	ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserDetail, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) PassLogin(ctx context.Context, in *UserPassLoginReq, opts ...grpc.CallOption) (*UserPassLoginResp, error) {
	out := new(UserPassLoginResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/PassLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) OauthToken(ctx context.Context, in *OauthTokenReq, opts ...grpc.CallOption) (*OauthTokenResp, error) {
	out := new(OauthTokenResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/OauthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOauthToken(ctx context.Context, in *GetOauthTokenReq, opts ...grpc.CallOption) (*OauthTokenResp, error) {
	out := new(OauthTokenResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/GetOauthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateOauthAccount(ctx context.Context, in *CreateOauthAccountReq, opts ...grpc.CallOption) (*CreateOauthAccountResp, error) {
	out := new(CreateOauthAccountResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/CreateOauthAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetOauthAccountSecret(ctx context.Context, in *ResetOauthAccountReq, opts ...grpc.CallOption) (*CreateOauthAccountResp, error) {
	out := new(CreateOauthAccountResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/ResetOauthAccountSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOauthAccount(ctx context.Context, in *GetOauthAccountReq, opts ...grpc.CallOption) (*GetOauthAccountResp, error) {
	out := new(GetOauthAccountResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/GetOauthAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserResp, error) {
	out := new(ListUserResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error) {
	out := new(ListRoleResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/ListRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	out := new(DeleteUserResp)
	err := c.cc.Invoke(ctx, "/UserService.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/UserService.UserService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	PassLogin(context.Context, *UserPassLoginReq) (*UserPassLoginResp, error)
	OauthToken(context.Context, *OauthTokenReq) (*OauthTokenResp, error)
	GetOauthToken(context.Context, *GetOauthTokenReq) (*OauthTokenResp, error)
	CreateOauthAccount(context.Context, *CreateOauthAccountReq) (*CreateOauthAccountResp, error)
	//重置oauth 的secret
	ResetOauthAccountSecret(context.Context, *ResetOauthAccountReq) (*CreateOauthAccountResp, error)
	//获取某个指定平台组织的oauth账号信息
	GetOauthAccount(context.Context, *GetOauthAccountReq) (*GetOauthAccountResp, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	ListUser(context.Context, *ListUserReq) (*ListUserResp, error)
	ListRole(context.Context, *ListRoleReq) (*ListRoleResp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error)
	UserInfo(context.Context, *UserInfoReq) (*UserDetail, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) PassLogin(context.Context, *UserPassLoginReq) (*UserPassLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassLogin not implemented")
}
func (UnimplementedUserServiceServer) OauthToken(context.Context, *OauthTokenReq) (*OauthTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OauthToken not implemented")
}
func (UnimplementedUserServiceServer) GetOauthToken(context.Context, *GetOauthTokenReq) (*OauthTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOauthToken not implemented")
}
func (UnimplementedUserServiceServer) CreateOauthAccount(context.Context, *CreateOauthAccountReq) (*CreateOauthAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOauthAccount not implemented")
}
func (UnimplementedUserServiceServer) ResetOauthAccountSecret(context.Context, *ResetOauthAccountReq) (*CreateOauthAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetOauthAccountSecret not implemented")
}
func (UnimplementedUserServiceServer) GetOauthAccount(context.Context, *GetOauthAccountReq) (*GetOauthAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOauthAccount not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) ListUser(context.Context, *ListUserReq) (*ListUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServiceServer) ListRole(context.Context, *ListRoleReq) (*ListRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) UserInfo(context.Context, *UserInfoReq) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_PassLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPassLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PassLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/PassLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PassLogin(ctx, req.(*UserPassLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_OauthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OauthTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).OauthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/OauthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).OauthToken(ctx, req.(*OauthTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOauthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOauthTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOauthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/GetOauthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOauthToken(ctx, req.(*GetOauthTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateOauthAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOauthAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateOauthAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/CreateOauthAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateOauthAccount(ctx, req.(*CreateOauthAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetOauthAccountSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetOauthAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetOauthAccountSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/ResetOauthAccountSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetOauthAccountSecret(ctx, req.(*ResetOauthAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOauthAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOauthAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOauthAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/GetOauthAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOauthAccount(ctx, req.(*GetOauthAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*ListUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/ListRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListRole(ctx, req.(*ListRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.UserService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PassLogin",
			Handler:    _UserService_PassLogin_Handler,
		},
		{
			MethodName: "OauthToken",
			Handler:    _UserService_OauthToken_Handler,
		},
		{
			MethodName: "GetOauthToken",
			Handler:    _UserService_GetOauthToken_Handler,
		},
		{
			MethodName: "CreateOauthAccount",
			Handler:    _UserService_CreateOauthAccount_Handler,
		},
		{
			MethodName: "ResetOauthAccountSecret",
			Handler:    _UserService_ResetOauthAccountSecret_Handler,
		},
		{
			MethodName: "GetOauthAccount",
			Handler:    _UserService_GetOauthAccount_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _UserService_ListRole_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserService.proto",
}
