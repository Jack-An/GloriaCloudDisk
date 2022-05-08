package main

import (
	"GloriaCloudDisk/common"
	"GloriaCloudDisk/user/api/internal/config"
	"GloriaCloudDisk/user/api/internal/handler"
	"GloriaCloudDisk/user/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *common.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusOK, common.NewDefaultError("unknown")
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
