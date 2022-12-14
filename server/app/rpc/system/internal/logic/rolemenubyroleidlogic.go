package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuByRoleIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleMenuByRoleIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuByRoleIDLogic {
	return &RoleMenuByRoleIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleMenuByRoleIDLogic) RoleMenuByRoleID(in *pb.RoleID) (*pb.RoleMenuResponse, error) {
	rolemenus, err := l.svcCtx.RoleMenuModel.FindByRoleID(l.ctx, l.svcCtx.Redis, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	pbRoleMeus := []*pb.RoleMenu{}
	for _, rolemenu := range rolemenus {
		pbRoleMeus = append(pbRoleMeus, &pb.RoleMenu{Id: rolemenu.RoleId, MenuID: rolemenu.MenuId, RoleID: rolemenu.RoleId})
	}

	return &pb.RoleMenuResponse{
		Rolemenus: pbRoleMeus,
	}, nil
}
