package main

import (
	"flag"
	"fmt"
	"github.com/wordpress-plus/api-app/doc/swagger"
	"github.com/wordpress-plus/api-app/internal/config"
	"github.com/wordpress-plus/api-app/internal/handler"
	"github.com/wordpress-plus/api-app/internal/middleware/gmw"
	"github.com/wordpress-plus/api-app/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile string

func init() {
	if con := os.Getenv("config"); len(con) != 0 {
		configFile = con
	} else {
		configFile = "etc/app-api-local.yaml"
	}
	fmt.Println("use config: " + configFile)
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svc.SvcCtx = svc.NewServiceContext(c)

	// mw: logger
	// server.Use(gmw.NewAddLogMiddleware(ctx).Handle)
	server.Use(gmw.NewAuthMiddleware(svc.SvcCtx).Handle)
	server.Use(gmw.NewRecordOpsMiddleware(svc.SvcCtx).Handle)

	handler.RegisterHandlers(server, svc.SvcCtx)
	swagger.RegisterSwagger(c.Mode, server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
