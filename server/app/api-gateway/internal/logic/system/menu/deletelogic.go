package menu

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

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

func (l *DeleteLogic) Delete(req *types.MenuDeleteRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.MenuID{ID: req.ID}
	_, err = l.svcCtx.SystemRpcClient.DeleteMenu_RoleMenu_UserMenuParam(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == "存在子菜单不可删除" {
			return nil, errorx.New(err, "存在子菜单不可删除").WithMeta("SystemRpcClient.DeleteMenu_RoleMenu_UserMenuParam", err.Error(), param)
		}
		return nil, errorx.New(err, "删除菜单失败").WithMeta("SystemRpcClient.DeleteMenu_RoleMenu_UserMenuParam", err.Error(), param)
	}

	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
