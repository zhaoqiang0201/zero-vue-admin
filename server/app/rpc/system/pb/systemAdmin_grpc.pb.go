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
	UserDetail(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	UserDetailByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error)
	UserPaging(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*UserPagingResponse, error)
	UserTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error)
	CreateUser_UserRole(ctx context.Context, in *CreateUser_UserRoleRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteSoftUser(ctx context.Context, in *DeleteSoftUserRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Empty, error)
	UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Empty, error)
	UserPageSet(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPageSetResponse, error)
	UpdateUserPageSet(ctx context.Context, in *UpdateUserPageSetRequest, opts ...grpc.CallOption) (*Empty, error)
	UserMenuParams(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserMenuParamsResponse, error)
	UserAllMenuParams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserMenuParamsResponse, error)
	UserRoleByUserID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserRoleResponse, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*Empty, error)
	RoleDetail(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*Role, error)
	RoleAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RoleAllResponse, error)
	MenuDetail(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*Menu, error)
	MenuAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MenuAllResponse, error)
	CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Empty, error)
	RoleMenuByRoleID(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleMenuResponse, error)
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

func (c *systemServiceClient) UserDetail(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserDetailByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserDetailByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserPaging(ctx context.Context, in *UserPagingRequest, opts ...grpc.CallOption) (*UserPagingResponse, error) {
	out := new(UserPagingResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserPaging", in, out, opts...)
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

func (c *systemServiceClient) CreateUser_UserRole(ctx context.Context, in *CreateUser_UserRoleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/CreateUser_UserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) DeleteSoftUser(ctx context.Context, in *DeleteSoftUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/DeleteSoftUser", in, out, opts...)
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

func (c *systemServiceClient) UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserPageSet(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPageSetResponse, error) {
	out := new(UserPageSetResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserPageSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateUserPageSet(ctx context.Context, in *UpdateUserPageSetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UpdateUserPageSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserMenuParams(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserMenuParamsResponse, error) {
	out := new(UserMenuParamsResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserMenuParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserAllMenuParams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserMenuParamsResponse, error) {
	out := new(UserMenuParamsResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserAllMenuParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UserRoleByUserID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserRoleResponse, error) {
	out := new(UserRoleResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UserRoleByUserID", in, out, opts...)
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

func (c *systemServiceClient) RoleDetail(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/pb.SystemService/RoleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RoleAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RoleAllResponse, error) {
	out := new(RoleAllResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/RoleAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) MenuDetail(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, "/pb.SystemService/MenuDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) MenuAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MenuAllResponse, error) {
	out := new(MenuAllResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/MenuAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/CreateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.SystemService/UpdateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RoleMenuByRoleID(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleMenuResponse, error) {
	out := new(RoleMenuResponse)
	err := c.cc.Invoke(ctx, "/pb.SystemService/RoleMenuByRoleID", in, out, opts...)
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
	UserDetail(context.Context, *UserID) (*User, error)
	UserDetailByName(context.Context, *UserName) (*User, error)
	UserPaging(context.Context, *UserPagingRequest) (*UserPagingResponse, error)
	UserTotal(context.Context, *Empty) (*Total, error)
	CreateUser_UserRole(context.Context, *CreateUser_UserRoleRequest) (*Empty, error)
	DeleteSoftUser(context.Context, *DeleteSoftUserRequest) (*Empty, error)
	DeleteUser(context.Context, *UserID) (*Empty, error)
	UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*Empty, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Empty, error)
	UserPageSet(context.Context, *UserID) (*UserPageSetResponse, error)
	UpdateUserPageSet(context.Context, *UpdateUserPageSetRequest) (*Empty, error)
	UserMenuParams(context.Context, *UserID) (*UserMenuParamsResponse, error)
	UserAllMenuParams(context.Context, *Empty) (*UserMenuParamsResponse, error)
	UserRoleByUserID(context.Context, *UserID) (*UserRoleResponse, error)
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*Empty, error)
	RoleDetail(context.Context, *RoleID) (*Role, error)
	RoleAll(context.Context, *Empty) (*RoleAllResponse, error)
	MenuDetail(context.Context, *MenuID) (*Menu, error)
	MenuAll(context.Context, *Empty) (*MenuAllResponse, error)
	CreateMenu(context.Context, *CreateMenuRequest) (*Empty, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*Empty, error)
	RoleMenuByRoleID(context.Context, *RoleID) (*RoleMenuResponse, error)
	Test(context.Context, *Empty) (*Total, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSystemServiceServer) UserDetail(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetail not implemented")
}
func (UnimplementedSystemServiceServer) UserDetailByName(context.Context, *UserName) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailByName not implemented")
}
func (UnimplementedSystemServiceServer) UserPaging(context.Context, *UserPagingRequest) (*UserPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPaging not implemented")
}
func (UnimplementedSystemServiceServer) UserTotal(context.Context, *Empty) (*Total, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserTotal not implemented")
}
func (UnimplementedSystemServiceServer) CreateUser_UserRole(context.Context, *CreateUser_UserRoleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser_UserRole not implemented")
}
func (UnimplementedSystemServiceServer) DeleteSoftUser(context.Context, *DeleteSoftUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSoftUser not implemented")
}
func (UnimplementedSystemServiceServer) DeleteUser(context.Context, *UserID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedSystemServiceServer) UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (UnimplementedSystemServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedSystemServiceServer) UserPageSet(context.Context, *UserID) (*UserPageSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPageSet not implemented")
}
func (UnimplementedSystemServiceServer) UpdateUserPageSet(context.Context, *UpdateUserPageSetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPageSet not implemented")
}
func (UnimplementedSystemServiceServer) UserMenuParams(context.Context, *UserID) (*UserMenuParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserMenuParams not implemented")
}
func (UnimplementedSystemServiceServer) UserAllMenuParams(context.Context, *Empty) (*UserMenuParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAllMenuParams not implemented")
}
func (UnimplementedSystemServiceServer) UserRoleByUserID(context.Context, *UserID) (*UserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRoleByUserID not implemented")
}
func (UnimplementedSystemServiceServer) UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedSystemServiceServer) RoleDetail(context.Context, *RoleID) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDetail not implemented")
}
func (UnimplementedSystemServiceServer) RoleAll(context.Context, *Empty) (*RoleAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAll not implemented")
}
func (UnimplementedSystemServiceServer) MenuDetail(context.Context, *MenuID) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuDetail not implemented")
}
func (UnimplementedSystemServiceServer) MenuAll(context.Context, *Empty) (*MenuAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuAll not implemented")
}
func (UnimplementedSystemServiceServer) CreateMenu(context.Context, *CreateMenuRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedSystemServiceServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedSystemServiceServer) RoleMenuByRoleID(context.Context, *RoleID) (*RoleMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleMenuByRoleID not implemented")
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

func _SystemService_UserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserDetail(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserDetailByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserDetailByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserDetailByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserDetailByName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserPaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserPaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserPaging",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserPaging(ctx, req.(*UserPagingRequest))
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

func _SystemService_CreateUser_UserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser_UserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).CreateUser_UserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/CreateUser_UserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).CreateUser_UserRole(ctx, req.(*CreateUser_UserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_DeleteSoftUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSoftUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).DeleteSoftUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/DeleteSoftUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).DeleteSoftUser(ctx, req.(*DeleteSoftUserRequest))
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

func _SystemService_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateUserPassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserPageSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserPageSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserPageSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserPageSet(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateUserPageSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPageSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateUserPageSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UpdateUserPageSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateUserPageSet(ctx, req.(*UpdateUserPageSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserMenuParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserMenuParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserMenuParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserMenuParams(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserAllMenuParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserAllMenuParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserAllMenuParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserAllMenuParams(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UserRoleByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UserRoleByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UserRoleByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UserRoleByUserID(ctx, req.(*UserID))
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

func _SystemService_RoleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RoleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/RoleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RoleDetail(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RoleAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RoleAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/RoleAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RoleAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_MenuDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).MenuDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/MenuDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).MenuDetail(ctx, req.(*MenuID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_MenuAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).MenuAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/MenuAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).MenuAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/CreateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).CreateMenu(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/UpdateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RoleMenuByRoleID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RoleMenuByRoleID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SystemService/RoleMenuByRoleID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RoleMenuByRoleID(ctx, req.(*RoleID))
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
			MethodName: "UserDetail",
			Handler:    _SystemService_UserDetail_Handler,
		},
		{
			MethodName: "UserDetailByName",
			Handler:    _SystemService_UserDetailByName_Handler,
		},
		{
			MethodName: "UserPaging",
			Handler:    _SystemService_UserPaging_Handler,
		},
		{
			MethodName: "UserTotal",
			Handler:    _SystemService_UserTotal_Handler,
		},
		{
			MethodName: "CreateUser_UserRole",
			Handler:    _SystemService_CreateUser_UserRole_Handler,
		},
		{
			MethodName: "DeleteSoftUser",
			Handler:    _SystemService_DeleteSoftUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _SystemService_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _SystemService_UpdateUserPassword_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _SystemService_UpdateUser_Handler,
		},
		{
			MethodName: "UserPageSet",
			Handler:    _SystemService_UserPageSet_Handler,
		},
		{
			MethodName: "UpdateUserPageSet",
			Handler:    _SystemService_UpdateUserPageSet_Handler,
		},
		{
			MethodName: "UserMenuParams",
			Handler:    _SystemService_UserMenuParams_Handler,
		},
		{
			MethodName: "UserAllMenuParams",
			Handler:    _SystemService_UserAllMenuParams_Handler,
		},
		{
			MethodName: "UserRoleByUserID",
			Handler:    _SystemService_UserRoleByUserID_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _SystemService_UpdateUserRole_Handler,
		},
		{
			MethodName: "RoleDetail",
			Handler:    _SystemService_RoleDetail_Handler,
		},
		{
			MethodName: "RoleAll",
			Handler:    _SystemService_RoleAll_Handler,
		},
		{
			MethodName: "MenuDetail",
			Handler:    _SystemService_MenuDetail_Handler,
		},
		{
			MethodName: "MenuAll",
			Handler:    _SystemService_MenuAll_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _SystemService_CreateMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _SystemService_UpdateMenu_Handler,
		},
		{
			MethodName: "RoleMenuByRoleID",
			Handler:    _SystemService_RoleMenuByRoleID_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _SystemService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doc/systemAdmin.proto",
}
