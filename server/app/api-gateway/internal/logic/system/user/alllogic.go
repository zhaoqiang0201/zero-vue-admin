package user

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllLogic {
	return &AllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllLogic) All() (resp *types.UserAllResponse, err error) {
	pagingUser, err := l.svcCtx.SystemRpcClient.UserPaging(l.ctx, &systemservice.UserPagingRequest{Page: 0, PageSize: 0})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Message() == sql.ErrNoRows.Error() {
			return &types.UserAllResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
				Total:              0,
				List:               []types.User{},
			}, nil
		}
		return nil, errorx.New(err, "获取分页用户失败").WithMeta("SystemRpcClient.PagingUserList", err.Error(), "")
	}
	state := map[bool]string{
		true:  "deleted",
		false: "resume",
	}
	list := []types.User{}
	for _, user := range pagingUser.Users {
		ut := types.User{
			ID:         user.ID,
			Name:       user.Name,
			NickName:   user.NickName,
			PassWord:   user.PassWord,
			UserType:   user.UserType,
			Email:      user.Email,
			Phone:      user.Phone,
			Department: user.Department,
			Position:   user.Position,
			CreateBy:   user.CreateBy,
			CreateTime: user.CreateTime,
			UpdateBy:   user.UpdateBy,
			UpdateTime: user.UpdateTime,
			DeleteBy:   user.DeleteBy,
			DeleteTime: user.DeleteTime,
			State:      state[user.DeleteTime != 0],
		}
		list = append(list, ut)
	}

	return &types.UserAllResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              int64(len(pagingUser.Users)),
		List:               list,
	}, nil
}
