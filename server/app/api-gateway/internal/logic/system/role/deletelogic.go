package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.RoleDeleteRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.RoleID{ID: req.ID}
	_, err = l.svcCtx.SystemRpcClient.DeleteRole(l.ctx, param)
	if err != nil {
		return nil, errorx.New(err, "删除角色失败").WithMeta("SystemRpcClient.DeleteRole", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
