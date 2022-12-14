package logic

import (
	"context"
	"database/sql"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSoftUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSoftUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSoftUserLogic {
	return &DeleteSoftUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSoftUserLogic) DeleteSoftUser(in *pb.DeleteSoftUserRequest) (*pb.Empty, error) {
	switch in.State {
	case "deleted":
		if err := l.svcCtx.UserModel.UpdateDeleteColumn(l.ctx, in.UserID, in.DeleteBy, sql.NullTime{Time: time.Now(), Valid: true}); err != nil {
			return nil, err
		}
	case "resume":
		if err := l.svcCtx.UserModel.UpdateDeleteColumn(l.ctx, in.UserID, in.DeleteBy, sql.NullTime{}); err != nil {
			return nil, err
		}
	}
	return &pb.Empty{}, nil
}
