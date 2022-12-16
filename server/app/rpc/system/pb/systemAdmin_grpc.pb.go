// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: doc/systemAdmin.proto

package pb

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

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UserInfoByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error)
	UserInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	RoleInfo(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*Role, error)
	MenuInfo(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*Menu, error)
	UserPageSetInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPageSet, error)
	SetUserPageSet(ctx context.Context, in *SetUserPageSetRequest, opts ...grpc.CallOption) (*Empty, error)
	AllRoleList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllRoleListResponse, error)
	PagingUserList(ctx context.Context, in *PagingRequest, opts ...grpc.CallOption) (*PagingUserListResponse, error)
	UserTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error)
	GetUserRoleByUserID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserRoleList, error)
	GetRoleMenuByRoleID(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleMenuList, error)
	GetUserMenuParams(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserMenuParamsList, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*Empty, error)
	SoftDeleteUser(ctx context.Context, in *SoftDeleteUserRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Empty, error)
	AddUserAndUserRole(ctx context.Context, in *AddUserAndUserRoleRequest, opts ...grpc.CallOption) (*Empty, error)
	Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserInfoByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserInfoByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RoleInfo(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/pb.SystemService/RoleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) MenuInfo(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, "/pb.SystemService/MenuInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserPageSetInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPageSet, error) {
	out := new(UserPageSet)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserPageSetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) SetUserPageSet(ctx context.Context, in *SetUserPageSetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/SetUserPageSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) AllRoleList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllRoleListResponse, error) {
	out := new(AllRoleListResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/AllRoleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) PagingUserList(ctx context.Context, in *PagingRequest, opts ...grpc.CallOption) (*PagingUserListResponse, error) {
	out := new(PagingUserListResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/PagingUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error) {
	out := new(Total)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserTotal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetUserRoleByUserID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserRoleList, error) {
	out := new(UserRoleList)
	err := c.cc.Invoke(ctx, "/pb.SystemService/GetUserRoleByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetRoleMenuByRoleID(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleMenuList, error) {
	out := new(RoleMenuList)
	err := c.cc.Invoke(ctx, "/pb.SystemService/GetRoleMenuByRoleID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetUserMenuParams(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserMenuParamsList, error) {
	out := new(UserMenuParamsList)
	err := c.cc.Invoke(ctx, "/pb.SystemService/GetUserMenuParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) SoftDeleteUser(ctx context.Context, in *SoftDeleteUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/SoftDeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) DeleteUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) AddUserAndUserRole(ctx context.Context, in *AddUserAndUserRoleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/AddUserAndUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error) {
	out := new(Total)
	err := c.cc.Invoke(ctx, "/pb.SystemService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
// All implementations must embed UnimplementedSystemServiceServer
// for forward compatibility
type SystemServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UserInfoByName(context.Context, *UserName) (*User, error)
	UserInfo(context.Context, *UserID) (*User, error)
	RoleInfo(context.Context, *RoleID) (*Role, error)
	MenuInfo(context.Context, *MenuID) (*Menu, error)
	UserPageSetInfo(context.Context, *UserID) (*UserPageSet, error)
	SetUserPageSet(context.Context, *SetUserPageSetRequest) (*Empty, error)
	AllRoleList(context.Context, *Empty) (*AllRoleListResponse, error)
	PagingUserList(context.Context, *PagingRequest) (*PagingUserListResponse, error)
	UserTotal(context.Context, *Empty) (*Total, error)
	GetUserRoleByUserID(context.Context, *UserID) (*UserRoleList, error)
	GetRoleMenuByRoleID(context.Context, *RoleID) (*RoleMenuList, error)
	GetUserMenuParams(context.Context, *UserID) (*UserMenuParamsList, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*Empty, error)
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*Empty, error)
	SoftDeleteUser(context.Context, *SoftDeleteUserRequest) (*Empty, error)
	DeleteUser(context.Context, *UserID) (*Empty, error)
	AddUserAndUserRole(context.Context, *AddUserAndUserRoleRequest) (*Empty, error)
	Test(context.Context, *Empty) (*Total, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSystemServiceServer) UserInfoByName(context.Context, *UserName) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoByName not implemented")
}
func (UnimplementedSystemServiceServer) UserInfo(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSystemServiceServer) RoleInfo(context.Context, *RoleID) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleInfo not implemented")
}
func (UnimplementedSystemServiceServer) MenuInfo(context.Context, *MenuID) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuInfo not implemented")
}
func (UnimplementedSystemServiceServer) UserPageSetInfo(context.Context, *UserID) (*UserPageSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPageSetInfo not implemented")
}
func (UnimplementedSystemServiceServer) SetUserPageSet(context.Context, *SetUserPageSetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPageSet not implemented")
}
func (UnimplementedSystemServiceServer) AllRoleList(context.Context, *Empty) (*AllRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllRoleList not implemented")
}
func (UnimplementedSystemServiceServer) PagingUserList(context.Context, *PagingRequest) (*PagingUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PagingUserList not implemented")
}
func (UnimplementedSystemServiceServer) UserTotal(context.Context, *Empty) (*Total, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserTotal not implemented")
}
func (UnimplementedSystemServiceServer) GetUserRoleByUserID(context.Context, *UserID) (*UserRoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoleByUserID not implemented")
}
func (UnimplementedSystemServiceServer) GetRoleMenuByRoleID(context.Context, *RoleID) (*RoleMenuList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleMenuByRoleID not implemented")
}
func (UnimplementedSystemServiceServer) GetUserMenuParams(context.Context, *UserID) (*UserMenuParamsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMenuParams not implemented")
}
func (UnimplementedSystemServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedSystemServiceServer) UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedSystemServiceServer) SoftDeleteUser(context.Context, *SoftDeleteUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDeleteUser not implemented")
}
func (UnimplementedSystemServiceServer) DeleteUser(context.Context, *UserID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedSystemServiceServer) AddUserAndUserRole(context.Context, *AddUserAndUserRoleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAndUserRole not implemented")
}
func (UnimplementedSystemServiceServer) Test(context.Context, *Empty) (*Total, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedSystemServiceServer) mustEmbedUnimplementedSystemServiceServer() {}

// UnsafeSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServiceServer will
// result in compilation errors.
type UnsafeSystemServiceServer interface {
	mustEmbedUnimplementedSystemServiceServer()
}

func RegisterSystemServiceServer(s grpc.ServiceRegistrar, srv SystemServiceServer) {
	s.RegisterService(&SystemService_ServiceDesc, srv)
}

func _SystemService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserInfoByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserInfoByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserInfoByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserInfoByName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserInfo(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RoleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RoleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/RoleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RoleInfo(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_MenuInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).MenuInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/MenuInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).MenuInfo(ctx, req.(*MenuID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserPageSetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserPageSetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserPageSetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserPageSetInfo(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_SetUserPageSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPageSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).SetUserPageSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/SetUserPageSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).SetUserPageSet(ctx, req.(*SetUserPageSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_AllRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).AllRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/AllRoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).AllRoleList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_PagingUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).PagingUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/PagingUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).PagingUserList(ctx, req.(*PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserTotal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserTotal(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetUserRoleByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetUserRoleByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/GetUserRoleByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetUserRoleByUserID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetRoleMenuByRoleID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetRoleMenuByRoleID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/GetRoleMenuByRoleID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetRoleMenuByRoleID(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetUserMenuParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetUserMenuParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/GetUserMenuParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetUserMenuParams(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateUserRole(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_SoftDeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoftDeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).SoftDeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/SoftDeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).SoftDeleteUser(ctx, req.(*SoftDeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).DeleteUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_AddUserAndUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAndUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).AddUserAndUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/AddUserAndUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).AddUserAndUserRole(ctx, req.(*AddUserAndUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).Test(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemService_ServiceDesc is the grpc.ServiceDesc for SystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SystemService_Login_Handler,
		},
		{
			MethodName: "UserInfoByName",
			Handler:    _SystemService_UserInfoByName_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _SystemService_UserInfo_Handler,
		},
		{
			MethodName: "RoleInfo",
			Handler:    _SystemService_RoleInfo_Handler,
		},
		{
			MethodName: "MenuInfo",
			Handler:    _SystemService_MenuInfo_Handler,
		},
		{
			MethodName: "UserPageSetInfo",
			Handler:    _SystemService_UserPageSetInfo_Handler,
		},
		{
			MethodName: "SetUserPageSet",
			Handler:    _SystemService_SetUserPageSet_Handler,
		},
		{
			MethodName: "AllRoleList",
			Handler:    _SystemService_AllRoleList_Handler,
		},
		{
			MethodName: "PagingUserList",
			Handler:    _SystemService_PagingUserList_Handler,
		},
		{
			MethodName: "UserTotal",
			Handler:    _SystemService_UserTotal_Handler,
		},
		{
			MethodName: "GetUserRoleByUserID",
			Handler:    _SystemService_GetUserRoleByUserID_Handler,
		},
		{
			MethodName: "GetRoleMenuByRoleID",
			Handler:    _SystemService_GetRoleMenuByRoleID_Handler,
		},
		{
			MethodName: "GetUserMenuParams",
			Handler:    _SystemService_GetUserMenuParams_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _SystemService_ChangePassword_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _SystemService_UpdateUserRole_Handler,
		},
		{
			MethodName: "SoftDeleteUser",
			Handler:    _SystemService_SoftDeleteUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _SystemService_DeleteUser_Handler,
		},
		{
			MethodName: "AddUserAndUserRole",
			Handler:    _SystemService_AddUserAndUserRole_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _SystemService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doc/systemAdmin.proto",
}
