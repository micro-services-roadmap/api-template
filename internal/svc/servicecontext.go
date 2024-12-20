package svc

import (
	"github.com/wordpress-plus/api-app/internal/config"
	"github.com/wordpress-plus/api-app/internal/middleware"
	"github.com/wordpress-plus/rpc-tracing/client/tracingipservice"
	"github.com/wordpress-plus/rpc-tracing/client/tracingviewservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var SvcCtx *ServiceContext

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware

	// tracing 模块
	TracingViewService tracingviewservice.TracingViewService
	TracingIPService   tracingipservice.TracingIPService
}

func NewServiceContext(c config.Config) *ServiceContext {

	trackingClient := zrpc.MustNewClient(c.TracingRpc)
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,

		TracingViewService: tracingviewservice.NewTracingViewService(trackingClient),
		TracingIPService:   tracingipservice.NewTracingIPService(trackingClient),
	}
}
