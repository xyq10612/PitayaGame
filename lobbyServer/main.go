package main

import (
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"lobbyServer/service"
	"strings"
)

var app pitaya.Pitaya

func main() {
	serverType := "lobby"

	logrus.SetLevel(logrus.DebugLevel)

	config := config.NewDefaultBuilderConfig()
	builder := pitaya.NewDefaultBuilder(false, serverType, pitaya.Cluster, map[string]string{}, *config)

	app = builder.Build()

	defer app.Shutdown()

	initServices()

	app.Start()
}

func initServices() {
	account := service.NewAccountService(app)
	app.Register(account,
		component.WithName("account"),
		component.WithNameFunc(strings.ToLower))
	app.RegisterRemote(account,
		component.WithName("account"),
		component.WithNameFunc(strings.ToLower))
}
