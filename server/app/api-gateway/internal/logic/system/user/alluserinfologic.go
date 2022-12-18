package user

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUserInfoLogic {
	return &AllUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllUserInfoLogic) AllUserInfo() (resp *types.PagingUserResponse, err error) {
	userlist, err := l.svcCtx.SystemRpcClient.PagingUserList(l.ctx, &systemservice.PagingUserListRequest{Page: 0, PageSize: 0})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Message() == sql.ErrNoRows.Error() {
			return &types.PagingUserResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
				PagingCommonResponse: types.PagingCommonResponse{Page: 0, PageSize: 0, Total: 0},
				List:                 []types.User{},
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("PagingUserList", err.Error(), "")
	}
	state := map[bool]string{
		true:  "deleted",
		false: "resume",
	}
	list := []types.User{}
	for _, user := range userlist.List {
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

	return &types.PagingUserResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
		PagingCommonResponse: types.PagingCommonResponse{Page: 0, PageSize: 0, Total: int64(len(userlist.List))},
		List:                 list,
	}, nil
}
