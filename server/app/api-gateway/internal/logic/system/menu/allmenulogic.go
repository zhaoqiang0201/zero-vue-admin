package menu

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllMenuLogic {
	return &AllMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllMenuLogic) AllMenu() (resp *types.AllMenuResponse, err error) {
	var (
		menuList           []types.Menu
		userMenuParamsList = []types.Parameter{}
		msgErrList         = errorx.MsgErrList{}
	)

	param := &systemservice.Empty{}
	allmenus, err := l.svcCtx.SystemRpcClient.AllMenuList(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.AllMenuResponse{HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"}, Total: 0, List: []types.Menu{}}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.AllMenuList", err.Error(), param)
	}
	for _, menu := range allmenus.List {
		m := types.Menu{
			ID:        menu.ID,
			ParentId:  menu.ParentID,
			Name:      menu.Name,
			Path:      menu.Path,
			Component: menu.Component,
			Sort:      menu.Sort,
			MenuMeta:  types.MenuMeta{Icon: menu.Icon, Title: menu.Title},
		}
		menuList = append(menuList, m)
	}

	usermenuparam, err := l.svcCtx.SystemRpcClient.AllUserMenuParams(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {

		} else {
			msgErrList.WithMeta("SystemRpcClient.AllUserMenuParams", err.Error(), param)
		}
		usermenuparam = new(systemservice.UserMenuParamsList)
	}
	for _, v := range usermenuparam.UserMenuParams {
		p := types.Parameter{
			ID:     v.ID,
			UserID: v.UserID,
			MenuID: v.MenuID,
			Type:   v.Type,
			Key:    v.Key,
			Value:  v.Value,
		}
		userMenuParamsList = append(userMenuParamsList, p)
	}

	menuTree := genMenuTreeMap(menuList, userMenuParamsList)
	var (
		msg   string = "OK"
		count int    = len(msgErrList.List)
	)
	if count != 0 {
		msg = fmt.Sprintf("Not OK(%d)", count)
	}
	return &types.AllMenuResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList},
		Total:              int64(len(allmenus.List)),
		List:               menuTree,
	}, nil
}
