package gmw

import (
	"context"
	"github.com/micro-services-roadmap/oneid-core/modelo"
	"github.com/micro-services-roadmap/oneid-core/utilo"
	"github.com/wordpress-plus/app-api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AuthMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
	return &AuthMiddleware{svcCtx: svcCtx}
}

// Handle parse jwt and add value to context
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claims, err := utilo.DecodeJwt(r.Header.Get("Authorization"))
		if err != nil || claims.Value == nil {
			ctx := context.WithValue(r.Context(), "userName", "admin-admin")
			ctx = context.WithValue(ctx, "uid", int64(-1))
			next(w, r.WithContext(ctx))
			return
		}

		if jwtUser, err := modelo.UserUnMarshal(*claims.Value); err != nil {
			logx.Errorf("[NewAuthMiddleware]jwt user unmarshal error: %v", err)
			next(w, r)
		} else {
			ctx := context.WithValue(r.Context(), modelo.Name, jwtUser.Name)
			ctx = context.WithValue(ctx, modelo.ID, jwtUser.Id)
			ctx = context.WithValue(ctx, modelo.Email, jwtUser.Email)
			next(w, r.WithContext(ctx))
		}
	}
}
