package actuator

import (
	"context"

	"github.com/wordpress-plus/app-api/internal/svc"
	"github.com/wordpress-plus/app-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() (resp *types.HealthCheckResp, err error) {
	return &types.HealthCheckResp{Appstatus: "UP"}, nil
}
