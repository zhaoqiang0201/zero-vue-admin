// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemmenu "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/menu"
	systemrole "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/role"
	systemuser "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/user"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: systemuser.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/paginguser",
					Handler: systemuser.PagingUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/allrole",
					Handler: systemrole.AllRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/usermenus",
					Handler: systemmenu.UserMenusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/menu"),
	)
}
