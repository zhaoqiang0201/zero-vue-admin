// Code generated by goctl. DO NOT EDIT!
// Source: systemAdmin.proto

package server

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/logic"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"
)

type SystemServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSystemServiceServer
}

func NewSystemServiceServer(svcCtx *svc.ServiceContext) *SystemServiceServer {
	return &SystemServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *SystemServiceServer) UserInfoByName(ctx context.Context, in *pb.UserName) (*pb.User, error) {
	l := logic.NewUserInfoByNameLogic(ctx, s.svcCtx)
	return l.UserInfoByName(in)
}

func (s *SystemServiceServer) UserInfo(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *SystemServiceServer) RoleInfo(ctx context.Context, in *pb.RoleID) (*pb.Role, error) {
	l := logic.NewRoleInfoLogic(ctx, s.svcCtx)
	return l.RoleInfo(in)
}

func (s *SystemServiceServer) MenuInfo(ctx context.Context, in *pb.MenuID) (*pb.Menu, error) {
	l := logic.NewMenuInfoLogic(ctx, s.svcCtx)
	return l.MenuInfo(in)
}

func (s *SystemServiceServer) UserPageSetInfo(ctx context.Context, in *pb.UserID) (*pb.UserPageSet, error) {
	l := logic.NewUserPageSetInfoLogic(ctx, s.svcCtx)
	return l.UserPageSetInfo(in)
}

func (s *SystemServiceServer) SetUserPageSet(ctx context.Context, in *pb.SetUserPageSetRequest) (*pb.Empty, error) {
	l := logic.NewSetUserPageSetLogic(ctx, s.svcCtx)
	return l.SetUserPageSet(in)
}

func (s *SystemServiceServer) AllRoleList(ctx context.Context, in *pb.Empty) (*pb.AllRoleListResponse, error) {
	l := logic.NewAllRoleListLogic(ctx, s.svcCtx)
	return l.AllRoleList(in)
}

func (s *SystemServiceServer) PagingUserList(ctx context.Context, in *pb.PagingRequest) (*pb.PagingUserListResponse, error) {
	l := logic.NewPagingUserListLogic(ctx, s.svcCtx)
	return l.PagingUserList(in)
}

func (s *SystemServiceServer) UserTotal(ctx context.Context, in *pb.Empty) (*pb.Total, error) {
	l := logic.NewUserTotalLogic(ctx, s.svcCtx)
	return l.UserTotal(in)
}

func (s *SystemServiceServer) GetUserRoleByUserID(ctx context.Context, in *pb.UserID) (*pb.UserRoleList, error) {
	l := logic.NewGetUserRoleByUserIDLogic(ctx, s.svcCtx)
	return l.GetUserRoleByUserID(in)
}

func (s *SystemServiceServer) GetRoleMenuByRoleID(ctx context.Context, in *pb.RoleID) (*pb.RoleMenuList, error) {
	l := logic.NewGetRoleMenuByRoleIDLogic(ctx, s.svcCtx)
	return l.GetRoleMenuByRoleID(in)
}

func (s *SystemServiceServer) GetUserMenuParams(ctx context.Context, in *pb.UserID) (*pb.UserMenuParamsList, error) {
	l := logic.NewGetUserMenuParamsLogic(ctx, s.svcCtx)
	return l.GetUserMenuParams(in)
}

func (s *SystemServiceServer) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.Empty, error) {
	l := logic.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

func (s *SystemServiceServer) UpdateUserRole(ctx context.Context, in *pb.UpdateUserRoleRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserRole(in)
}

func (s *SystemServiceServer) SoftDeleteUser(ctx context.Context, in *pb.SoftDeleteUserRequest) (*pb.Empty, error) {
	l := logic.NewSoftDeleteUserLogic(ctx, s.svcCtx)
	return l.SoftDeleteUser(in)
}

func (s *SystemServiceServer) Test(ctx context.Context, in *pb.Empty) (*pb.Total, error) {
	l := logic.NewTestLogic(ctx, s.svcCtx)
	return l.Test(in)
}
