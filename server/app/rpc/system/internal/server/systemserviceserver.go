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

func (s *SystemServiceServer) CasbinEnforcer(ctx context.Context, in *pb.CasbinEnforceRequest) (*pb.CasbinEnforceResponse, error) {
	l := logic.NewCasbinEnforcerLogic(ctx, s.svcCtx)
	return l.CasbinEnforcer(in)
}

func (s *SystemServiceServer) RefreshCasbinPolicy(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	l := logic.NewRefreshCasbinPolicyLogic(ctx, s.svcCtx)
	return l.RefreshCasbinPolicy(in)
}

func (s *SystemServiceServer) UserDetail(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	l := logic.NewUserDetailLogic(ctx, s.svcCtx)
	return l.UserDetail(in)
}

func (s *SystemServiceServer) UserDetailByName(ctx context.Context, in *pb.UserName) (*pb.User, error) {
	l := logic.NewUserDetailByNameLogic(ctx, s.svcCtx)
	return l.UserDetailByName(in)
}

func (s *SystemServiceServer) UserPaging(ctx context.Context, in *pb.UserPagingRequest) (*pb.UserPagingResponse, error) {
	l := logic.NewUserPagingLogic(ctx, s.svcCtx)
	return l.UserPaging(in)
}

func (s *SystemServiceServer) UserTotal(ctx context.Context, in *pb.Empty) (*pb.Total, error) {
	l := logic.NewUserTotalLogic(ctx, s.svcCtx)
	return l.UserTotal(in)
}

func (s *SystemServiceServer) CreateUser_UserRole(ctx context.Context, in *pb.CreateUser_UserRoleRequest) (*pb.Empty, error) {
	l := logic.NewCreateUserUserRoleLogic(ctx, s.svcCtx)
	return l.CreateUser_UserRole(in)
}

func (s *SystemServiceServer) DeleteSoftUser(ctx context.Context, in *pb.DeleteSoftUserRequest) (*pb.Empty, error) {
	l := logic.NewDeleteSoftUserLogic(ctx, s.svcCtx)
	return l.DeleteSoftUser(in)
}

func (s *SystemServiceServer) DeleteUser(ctx context.Context, in *pb.UserID) (*pb.Empty, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *SystemServiceServer) UpdateUserPassword(ctx context.Context, in *pb.UpdateUserPasswordRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserPasswordLogic(ctx, s.svcCtx)
	return l.UpdateUserPassword(in)
}

func (s *SystemServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *SystemServiceServer) UpdateUserCurrentRole(ctx context.Context, in *pb.UpdateUserCurrentRoleRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserCurrentRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserCurrentRole(in)
}

func (s *SystemServiceServer) UserPageSet(ctx context.Context, in *pb.UserID) (*pb.UserPageSetResponse, error) {
	l := logic.NewUserPageSetLogic(ctx, s.svcCtx)
	return l.UserPageSet(in)
}

func (s *SystemServiceServer) UpdateUserPageSet(ctx context.Context, in *pb.UpdateUserPageSetRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserPageSetLogic(ctx, s.svcCtx)
	return l.UpdateUserPageSet(in)
}

func (s *SystemServiceServer) UserMenuParamsByUserID(ctx context.Context, in *pb.UserID) (*pb.UserMenuParamsResponse, error) {
	l := logic.NewUserMenuParamsByUserIDLogic(ctx, s.svcCtx)
	return l.UserMenuParamsByUserID(in)
}

func (s *SystemServiceServer) UserAllMenuParams(ctx context.Context, in *pb.Empty) (*pb.UserMenuParamsResponse, error) {
	l := logic.NewUserAllMenuParamsLogic(ctx, s.svcCtx)
	return l.UserAllMenuParams(in)
}

func (s *SystemServiceServer) UpdateUserMenuParams(ctx context.Context, in *pb.UpdateUserMenuParamsRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserMenuParamsLogic(ctx, s.svcCtx)
	return l.UpdateUserMenuParams(in)
}

func (s *SystemServiceServer) UserRoleByUserID(ctx context.Context, in *pb.UserID) (*pb.UserRoleResponse, error) {
	l := logic.NewUserRoleByUserIDLogic(ctx, s.svcCtx)
	return l.UserRoleByUserID(in)
}

func (s *SystemServiceServer) UpdateUserRole(ctx context.Context, in *pb.UpdateUserRoleRequest) (*pb.Empty, error) {
	l := logic.NewUpdateUserRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserRole(in)
}

func (s *SystemServiceServer) RoleDetail(ctx context.Context, in *pb.RoleID) (*pb.Role, error) {
	l := logic.NewRoleDetailLogic(ctx, s.svcCtx)
	return l.RoleDetail(in)
}

func (s *SystemServiceServer) RoleAll(ctx context.Context, in *pb.Empty) (*pb.RoleAllResponse, error) {
	l := logic.NewRoleAllLogic(ctx, s.svcCtx)
	return l.RoleAll(in)
}

func (s *SystemServiceServer) DeleteRole(ctx context.Context, in *pb.RoleID) (*pb.Empty, error) {
	l := logic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

func (s *SystemServiceServer) DeleteSoftRole(ctx context.Context, in *pb.DeleteSoftRoleRequest) (*pb.Empty, error) {
	l := logic.NewDeleteSoftRoleLogic(ctx, s.svcCtx)
	return l.DeleteSoftRole(in)
}

func (s *SystemServiceServer) CreateRole(ctx context.Context, in *pb.CreateRoleRequest) (*pb.Empty, error) {
	l := logic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *SystemServiceServer) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.Empty, error) {
	l := logic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *SystemServiceServer) MenuDetail(ctx context.Context, in *pb.MenuID) (*pb.Menu, error) {
	l := logic.NewMenuDetailLogic(ctx, s.svcCtx)
	return l.MenuDetail(in)
}

func (s *SystemServiceServer) MenuAll(ctx context.Context, in *pb.Empty) (*pb.MenuAllResponse, error) {
	l := logic.NewMenuAllLogic(ctx, s.svcCtx)
	return l.MenuAll(in)
}

func (s *SystemServiceServer) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.Empty, error) {
	l := logic.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

func (s *SystemServiceServer) UpdateMenu(ctx context.Context, in *pb.UpdateMenuRequest) (*pb.Empty, error) {
	l := logic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *SystemServiceServer) DeleteMenu_RoleMenu_UserMenuParam(ctx context.Context, in *pb.MenuID) (*pb.Empty, error) {
	l := logic.NewDeleteMenuRoleMenuUserMenuParamLogic(ctx, s.svcCtx)
	return l.DeleteMenu_RoleMenu_UserMenuParam(in)
}

func (s *SystemServiceServer) RoleMenuByRoleID(ctx context.Context, in *pb.RoleID) (*pb.RoleMenuResponse, error) {
	l := logic.NewRoleMenuByRoleIDLogic(ctx, s.svcCtx)
	return l.RoleMenuByRoleID(in)
}

func (s *SystemServiceServer) UpdateRoleMenus(ctx context.Context, in *pb.UpdateRoleMenusRequest) (*pb.Empty, error) {
	l := logic.NewUpdateRoleMenusLogic(ctx, s.svcCtx)
	return l.UpdateRoleMenus(in)
}

func (s *SystemServiceServer) APIDetail(ctx context.Context, in *pb.ApiID) (*pb.API, error) {
	l := logic.NewAPIDetailLogic(ctx, s.svcCtx)
	return l.APIDetail(in)
}

func (s *SystemServiceServer) APIAll(ctx context.Context, in *pb.Empty) (*pb.APIAllResponse, error) {
	l := logic.NewAPIAllLogic(ctx, s.svcCtx)
	return l.APIAll(in)
}

func (s *SystemServiceServer) APIPaging(ctx context.Context, in *pb.APIPagingRequest) (*pb.APIPagingResponse, error) {
	l := logic.NewAPIPagingLogic(ctx, s.svcCtx)
	return l.APIPaging(in)
}

func (s *SystemServiceServer) APITotal(ctx context.Context, in *pb.APIPagingRequest) (*pb.Total, error) {
	l := logic.NewAPITotalLogic(ctx, s.svcCtx)
	return l.APITotal(in)
}

func (s *SystemServiceServer) CreateAPI(ctx context.Context, in *pb.CreateAPIRequest) (*pb.Empty, error) {
	l := logic.NewCreateAPILogic(ctx, s.svcCtx)
	return l.CreateAPI(in)
}

func (s *SystemServiceServer) UpdateAPI(ctx context.Context, in *pb.API) (*pb.Empty, error) {
	l := logic.NewUpdateAPILogic(ctx, s.svcCtx)
	return l.UpdateAPI(in)
}

func (s *SystemServiceServer) DeleteAPIAndCasbin(ctx context.Context, in *pb.DeleteAPIAndCasbinRequest) (*pb.Empty, error) {
	l := logic.NewDeleteAPIAndCasbinLogic(ctx, s.svcCtx)
	return l.DeleteAPIAndCasbin(in)
}

func (s *SystemServiceServer) DeleteAPIMultipleAndCasbin(ctx context.Context, in *pb.DeleteAPIMultipleAndCasbinRequest) (*pb.Empty, error) {
	l := logic.NewDeleteAPIMultipleAndCasbinLogic(ctx, s.svcCtx)
	return l.DeleteAPIMultipleAndCasbin(in)
}

func (s *SystemServiceServer) Test(ctx context.Context, in *pb.Empty) (*pb.Total, error) {
	l := logic.NewTestLogic(ctx, s.svcCtx)
	return l.Test(in)
}
