package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"proxyServer/service"
	"strings"
)

var app pitaya.Pitaya

func main() {
	serverType := "proxy"

	port := flag.Int("port", 40000, "the port to listen")
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	config := config.NewDefaultBuilderConfig()
	builder := pitaya.NewDefaultBuilder(true, serverType, pitaya.Cluster, map[string]string{}, *config)
	builder.AddAcceptor(newAcceptor(*port))

	app = builder.Build()

	defer app.Shutdown()

	initServices()

	app.Start()
}

func newAcceptor(port int) acceptor.Acceptor {
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	return tcp
}

func initServices() {
	account := service.NewAccountService(app)
	app.Register(account,
		component.WithName("account"),
		component.WithNameFunc(strings.ToLower))
}
