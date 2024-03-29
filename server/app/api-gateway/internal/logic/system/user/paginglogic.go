package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"sync"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingLogic {
	return &PagingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PagingLogic) Paging(req *types.UserPagingRequest) (resp *types.UserPagingResponse, err error) {
	var total int64
	var msgErrList = errorx.MsgErrList{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				logx.Error(r)
			}
		}()
		userTotalParam := &systemservice.Empty{}
		ptotal, err := l.svcCtx.SystemRpcClient.UserTotal(l.ctx, userTotalParam)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				msgErrList.WithMeta("SystemRpcClient.UserTotal", err.Error(), userTotalParam)
			} else {
				msgErrList.WithMeta("SystemRpcClient.UserTotal", err.Error(), userTotalParam)
			}
		}
		if ptotal != nil {
			total = ptotal.Total
		}
	}()

	pagingUserParam := &systemservice.UserPagingRequest{Page: req.Page, PageSize: req.PageSize, NameX: req.NameX}
	userlist, err := l.svcCtx.SystemRpcClient.UserPaging(l.ctx, pagingUserParam)
	if err != nil {
		e, _ := status.FromError(err)
		if e.Message() == sql.ErrNoRows.Error() {
			return &types.UserPagingResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
				PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
				List:                 []types.User{},
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UserPaging", err.Error(), pagingUserParam)
	}

	var (
		state = map[bool]string{
			true:  "deleted",
			false: "resume",
		}
	)
	list := []types.User{}
	for _, user := range userlist.Users {
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

	//获取用户角色
	mr.MapReduce(
		func(source chan<- interface{}) {
			for _, user := range list {
				source <- user.ID
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			userid := item.(uint64)
			userRoleByUserIDParam := &systemservice.UserID{ID: userid}
			userroles, err := l.svcCtx.SystemRpcClient.UserRoleByUserID(l.ctx, userRoleByUserIDParam)
			if err != nil {
				l.Error(err)
				msgErrList.WithMeta("SystemRpcClient.GetUserRoleByUserID", err.Error(), userRoleByUserIDParam)
				return
			}

			writer.Write(userroles)
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				userroles, ok := v.(*systemservice.UserRoleResponse)
				if ok {
					for _, userrole := range userroles.UserRoles {
						for index, user := range list {
							if user.ID == userrole.UserID {
								//roleallList = append(roleallList, userrole.RoleID)
								list[index].RoleList = append(list[index].RoleList, userrole.RoleID)
							}
						}
					}
				} else {
					logx.Errorf("mr reducer断言失败, 实际类型: (%T), 断言类型: (*systemservice.UserRoleList)", v, v)
				}

			}
		},
	)
	wg.Wait()
	var (
		msg     = "OK"
		elcount = len(msgErrList.List)
	)
	if elcount != 0 {
		msg = fmt.Sprintf("Not OK(%d), %v", elcount, msgErrList.List)
	}
	return &types.UserPagingResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
		List:                 list,
	}, nil
}
