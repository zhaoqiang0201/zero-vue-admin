// Code generated by goctl. DO NOT EDIT!
// Source: systemAdmin.proto

package systemservice

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AllRoleListResponse    = pb.AllRoleListResponse
	ChangePasswordRequest  = pb.ChangePasswordRequest
	Empty                  = pb.Empty
	LoginRequest           = pb.LoginRequest
	LoginResponse          = pb.LoginResponse
	Menu                   = pb.Menu
	MenuID                 = pb.MenuID
	PagingRequest          = pb.PagingRequest
	PagingUserListResponse = pb.PagingUserListResponse
	Role                   = pb.Role
	RoleID                 = pb.RoleID
	RoleMenu               = pb.RoleMenu
	RoleMenuList           = pb.RoleMenuList
	SetUserPageSetRequest  = pb.SetUserPageSetRequest
	SoftDeleteUserRequest  = pb.SoftDeleteUserRequest
	Total                  = pb.Total
	UpdateUserRoleRequest  = pb.UpdateUserRoleRequest
	User                   = pb.User
	UserID                 = pb.UserID
	UserMenuParams         = pb.UserMenuParams
	UserMenuParamsList     = pb.UserMenuParamsList
	UserName               = pb.UserName
	UserPageSet            = pb.UserPageSet
	UserRole               = pb.UserRole
	UserRoleList           = pb.UserRoleList

	SystemService interface {
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
		Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error)
	}

	defaultSystemService struct {
		cli zrpc.Client
	}
)

func NewSystemService(cli zrpc.Client) SystemService {
	return &defaultSystemService{
		cli: cli,
	}
}

func (m *defaultSystemService) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSystemService) UserInfoByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.UserInfoByName(ctx, in, opts...)
}

func (m *defaultSystemService) UserInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSystemService) RoleInfo(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*Role, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.RoleInfo(ctx, in, opts...)
}

func (m *defaultSystemService) MenuInfo(ctx context.Context, in *MenuID, opts ...grpc.CallOption) (*Menu, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.MenuInfo(ctx, in, opts...)
}

func (m *defaultSystemService) UserPageSetInfo(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserPageSet, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.UserPageSetInfo(ctx, in, opts...)
}

func (m *defaultSystemService) SetUserPageSet(ctx context.Context, in *SetUserPageSetRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.SetUserPageSet(ctx, in, opts...)
}

func (m *defaultSystemService) AllRoleList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllRoleListResponse, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.AllRoleList(ctx, in, opts...)
}

func (m *defaultSystemService) PagingUserList(ctx context.Context, in *PagingRequest, opts ...grpc.CallOption) (*PagingUserListResponse, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.PagingUserList(ctx, in, opts...)
}

func (m *defaultSystemService) UserTotal(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.UserTotal(ctx, in, opts...)
}

func (m *defaultSystemService) GetUserRoleByUserID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserRoleList, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.GetUserRoleByUserID(ctx, in, opts...)
}

func (m *defaultSystemService) GetRoleMenuByRoleID(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleMenuList, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.GetRoleMenuByRoleID(ctx, in, opts...)
}

func (m *defaultSystemService) GetUserMenuParams(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserMenuParamsList, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.GetUserMenuParams(ctx, in, opts...)
}

func (m *defaultSystemService) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

func (m *defaultSystemService) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in, opts...)
}

func (m *defaultSystemService) SoftDeleteUser(ctx context.Context, in *SoftDeleteUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.SoftDeleteUser(ctx, in, opts...)
}

func (m *defaultSystemService) Test(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Total, error) {
	client := pb.NewSystemServiceClient(m.cli.Conn())
	return client.Test(ctx, in, opts...)
}
