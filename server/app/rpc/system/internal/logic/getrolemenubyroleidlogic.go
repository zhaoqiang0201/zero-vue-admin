package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenuByRoleIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleMenuByRoleIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenuByRoleIDLogic {
	return &GetRoleMenuByRoleIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleMenuByRoleIDLogic) GetRoleMenuByRoleID(in *pb.RoleID) (*pb.RoleMenuList, error) {
	rolemenus, err := l.svcCtx.RoleMenuModel.FindByRoleID(l.ctx, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			if err == sqlc.ErrNotFound {
				return nil, errors.Wrapf(err, "无数据, 表: role_menu, 字段: user_id=%v", in.ID)
			}
			return nil, errors.Wrapf(err, "数据库查询失败, 表: role_menu, 字段: user_id=%v", in.ID)
		}
	}
	pbRoleMeus := []*pb.RoleMenu{}
	for _, rolemenu := range rolemenus {
		pbRoleMeus = append(pbRoleMeus, &pb.RoleMenu{Id: rolemenu.RoleId, MenuID: rolemenu.MenuId, RoleID: rolemenu.RoleId})
	}

	return &pb.RoleMenuList{
		Rolemenu: pbRoleMeus,
	}, nil
}
