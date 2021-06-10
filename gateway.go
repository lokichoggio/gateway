package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/lokichoggio/gateway/internal/common/errorx"
	"github.com/lokichoggio/gateway/internal/config"
	"github.com/lokichoggio/gateway/internal/handler"
	"github.com/lokichoggio/gateway/internal/middleware"
	"github.com/lokichoggio/gateway/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")
var listenOn = flag.Int("listen", 0, "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if *listenOn != 0 {
		c.RestConf.Port = *listenOn
	}

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	server.Use(middleware.ServiceMiddleware)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
