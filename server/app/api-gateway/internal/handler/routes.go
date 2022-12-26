// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemapi "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/api"
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
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/currentset",
					Handler: systemuser.CurrentUserSetingHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: systemuser.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: systemuser.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: systemuser.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemuser.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id/soft",
					Handler: systemuser.DeleteSoftHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id/role",
					Handler: systemuser.UpdateUserRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id/password",
					Handler: systemuser.UpdatePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: systemuser.AllHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: systemuser.PagingHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/page",
					Handler: systemuser.PageHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: systemrole.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: systemrole.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: systemrole.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id/soft",
					Handler: systemrole.DeleteSoftHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemrole.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: systemrole.AllHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/refreshpermission",
					Handler: systemrole.RefreshPermissionHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
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

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: systemmenu.AllHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: systemmenu.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemmenu.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: systemmenu.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: systemmenu.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id/userparam",
					Handler: systemmenu.UpdateUserParamHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: systemapi.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: systemapi.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/multiple",
					Handler: systemapi.DeleteMultipleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: systemapi.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemapi.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: systemapi.PagingHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: systemapi.AllHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/system/api"),
	)
}
