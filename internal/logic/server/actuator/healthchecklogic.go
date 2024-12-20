package actuator

import (
	"context"

	"github.com/wordpress-plus/api-app/internal/svc"
	"github.com/wordpress-plus/api-app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext, req *http.Request) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    req,
	}
}

func (l *HealthCheckLogic) HealthCheck() (*types.HealthCheckResp, error) {
	return &types.HealthCheckResp{Appstatus: "UP"}, nil
}
