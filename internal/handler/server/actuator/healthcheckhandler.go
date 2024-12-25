package actuator

import (
	"github.com/micro-services-roadmap/kit-common/gz"
	"github.com/micro-services-roadmap/oneid-core/modelo"
	"net/http"

	"github.com/wordpress-plus/app-api/internal/logic/server/actuator"
	"github.com/wordpress-plus/app-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := actuator.NewHealthCheckLogic(gz.ConvertHeaderMD(r), svcCtx)
		resp, err := l.HealthCheck()

		if err != nil {
			l.Logger.Errorf("Handler error: ", err)
			if v, ok := err.(*modelo.CodeError); ok {
				httpx.OkJsonCtx(r.Context(), w, modelo.Res(v.Code, nil, v.Msg))
			} else {
				httpx.OkJsonCtx(r.Context(), w, modelo.FailWithError(err))
			}
		} else {
			httpx.OkJsonCtx(r.Context(), w, modelo.OkWithData(resp))
		}
	}
}
