package main

import (
	"flag"
	"fmt"

	"activity/answer/api/internal/config"
	"activity/answer/api/internal/handler"
	"activity/answer/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/answer-api.yaml", "the config file")

var configFile1 = "answer/api/etc/answer-api.yaml"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(configFile1, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
