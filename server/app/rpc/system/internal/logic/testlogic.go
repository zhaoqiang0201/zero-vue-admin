package logic

import (
	"context"
	"database/sql"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestLogic) Test(in *pb.Empty) (*pb.Total, error) {
	total, err := l.svcCtx.UserModel.Total_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.Total{Total: 0}, err
		}
		return nil, err
	}
	return &pb.Total{Total: total}, nil
}
