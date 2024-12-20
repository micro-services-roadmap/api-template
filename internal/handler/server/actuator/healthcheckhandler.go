package actuator

import (
	"github.com/micro-services-roadmap/kit-common/errorx"
	"github.com/micro-services-roadmap/oneid-core/modelo"
	"net/http"

	"github.com/wordpress-plus/api-app/internal/logic/server/actuator"
	"github.com/wordpress-plus/api-app/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := actuator.NewHealthCheckLogic(r.Context(), svcCtx, r)
		resp, err := l.HealthCheck()

		if err != nil {
			l.Logger.Errorf("Handler error: ", err)
			err = errorx.GrpcError(err)
			if v, ok := err.(*modelo.CodeError); ok {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(v.Code, nil, v.Error()))
			} else {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, modelo.Res(888_888, nil, err.Error()))
			}
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
