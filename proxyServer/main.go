package main

import (
	"common/constants"
	"common/modules/db"
	"common/modules/db/redis"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"proxyServer/router"
	"proxyServer/service"
	"strings"
)

var app pitaya.Pitaya

func main() {
	serverType := constants.ProxyServer

	port := flag.Int("port", 40000, "the port to listen")
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	config := config.NewDefaultBuilderConfig()
	builder := pitaya.NewDefaultBuilder(true, serverType, pitaya.Cluster, map[string]string{}, *config)
	builder.AddAcceptor(newAcceptor(*port))

	app = builder.Build()
	pitaya.DefaultApp = app

	defer app.Shutdown()

	registerServices()
	registerModules()

	app.AddRoute(constants.LobbyServer, router.LobbyRouterFunc)

	app.Start()
}

func newAcceptor(port int) acceptor.Acceptor {
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	return tcp
}

func registerServices() {
	account := service.NewAccountService(app)
	app.Register(account,
		component.WithName("account"),
		component.WithNameFunc(strings.ToLower))
}

func registerModules() {
	r := redis.NewRedisStorage(redis.RedisConfig{
		Config: db.Config{
			Host:     "localhost",
			Port:     6379,
			Password: "fb123456",
		},
	})
	app.RegisterModule(r, constants.RedisModule)
}
