package conn

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/esManager/conn"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
)

func AllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := conn.NewAllLogic(r.Context(), svcCtx)
		resp, err := l.All()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
