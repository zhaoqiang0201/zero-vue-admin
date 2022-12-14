package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailByNameLogic {
	return &UserDetailByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailByNameLogic) UserDetailByName(in *pb.UserName) (*pb.User, error) {
	userinfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "查询数据库失败")
	}
	pUser := &pb.User{
		ID:         userinfo.Id,
		Name:       userinfo.Name,
		NickName:   userinfo.NickName,
		PassWord:   userinfo.Password,
		UserType:   userinfo.Type,
		Email:      userinfo.Email,
		Phone:      userinfo.Phone,
		Department: userinfo.Department,
		Position:   userinfo.Position,
		CreateBy:   userinfo.CreateBy,
		CreateTime: userinfo.CreateTime.Unix(),
		UpdateBy:   userinfo.UpdateBy,
		UpdateTime: userinfo.UpdateTime.Unix(),
		DeleteBy:   userinfo.DeleteBy,
	}
	if userinfo.DeleteTime.Valid {
		pUser.DeleteTime = userinfo.DeleteTime.Time.Unix()
	} else {
		pUser.DeleteTime = 0
	}
	return pUser, nil
}
