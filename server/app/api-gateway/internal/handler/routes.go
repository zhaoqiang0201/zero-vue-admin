// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	esManagerconn "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/esManager/conn"
	monitoringalarm "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/monitoring/alarm"
	monitoringalertrule "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/monitoring/alertrule"
	monitoringhosts "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/monitoring/hosts"
	monitoringstoreconnect "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/monitoring/store/connect"
	monitoringstorefile "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/monitoring/store/file"
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
					Path:    "/password",
					Handler: systemuser.UpdateLoginPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/page",
					Handler: systemuser.PageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/changerole",
					Handler: systemuser.ChangeRoleHandler(serverCtx),
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
					Path:    "/menupermission/:id",
					Handler: systemrole.MenuPermissionHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/menupermission/:id",
					Handler: systemrole.UpdateMenuPermissionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/apipermission/:id",
					Handler: systemrole.APIPermissionHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/apipermission/:id",
					Handler: systemrole.UpdateAPIPermissionHandler(serverCtx),
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
					Method:  http.MethodPut,
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

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: esManagerconn.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: esManagerconn.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: esManagerconn.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: esManagerconn.DELETEHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: esManagerconn.PagingHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/ping/:id",
					Handler: esManagerconn.PingHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/all",
					Handler: esManagerconn.AllHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/esmanager/conn"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: monitoringstoreconnect.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: monitoringstoreconnect.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: monitoringstoreconnect.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: monitoringstoreconnect.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: monitoringstoreconnect.PagingHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/monitoring/store/connect"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: monitoringstorefile.ListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/monitoring/store/file"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: monitoringalertrule.PagingHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/monitoring/alertrule"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseJWTToken, serverCtx.CheckUserExists, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/paging",
					Handler: monitoringhosts.PagingHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: monitoringhosts.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/",
					Handler: monitoringhosts.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/slience/all",
					Handler: monitoringhosts.GetAllSlienceJsonHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/slience/:host",
					Handler: monitoringhosts.GetSlienceHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/slience/:host",
					Handler: monitoringhosts.CreateUpdateSlienceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/slience",
					Handler: monitoringhosts.HandlerHostsSlienceHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/monitoring/hosts"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/webhook",
				Handler: monitoringalarm.WebhookHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/monitoring/alarm"),
	)
}
