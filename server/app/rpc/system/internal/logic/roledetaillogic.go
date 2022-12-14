package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDetailLogic {
	return &RoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDetailLogic) RoleDetail(in *pb.RoleID) (*pb.Role, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	res := &pb.Role{
		ID:         role.Id,
		Role:       role.Role,
		Name:       role.Name,
		CreateBy:   role.CreateBy,
		CreateTime: role.CreateTime.Unix(),
		UpdateBy:   role.UpdateBy,
		UpdateTime: role.UpdateTime.Unix(),
		DeleteBy:   role.DeleteBy,
	}
	if role.DeleteTime.Valid {
		res.DeleteTime = role.DeleteTime.Time.Unix()
	} else {
		res.DeleteTime = 0
	}
	return res, nil
}
