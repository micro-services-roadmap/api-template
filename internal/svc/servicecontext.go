package svc

import (
	"github.com/wordpress-plus/app-api/internal/config"
)

var SvcCtx *ServiceContext

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
