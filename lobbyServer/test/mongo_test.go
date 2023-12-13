package test

import (
	"common/constants"
	"common/modules/db"
	"common/modules/db/mongodb"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/helpers"
	"lobbyServer/model/accountModel"
	"testing"
	"time"
)

func mockApp() *pitaya.App {
	c := config.NewDefaultBuilderConfig()
	app := pitaya.NewDefaultApp(false, "testServer", pitaya.Cluster, map[string]string{}, *c).(*pitaya.App)

	return app
}

func TestAccountModel(t *testing.T) {
	app := mockApp()
	pitaya.DefaultApp = app

	mongo := mongodb.NewMongoStorage(mongodb.MongoConfig{
		Config: db.Config{
			Host:     "localhost",
			Port:     27017,
			Username: "debugeve",
			Password: "develop2023",
		},
		MaxPoolSize: 10,
	})
	app.RegisterModule(mongo, constants.MongoDBModule)

	go func() {
		app.Start()
	}()

	helpers.ShouldEventuallyReturn(t, func() bool {
		return app.IsRunning()
	}, true, time.Second, time.Second*10)

	name := fmt.Sprintf("test%s", time.Now().Format("20060102150405"))
	model := accountModel.AccountModel{
		Name:     name,
		Password: "pwd123456",
		Uid:      fmt.Sprintf("uid-%s", name),
	}
	err := model.New()
	assert.NoError(t, err)

	exist := accountModel.Exist(name)
	assert.Equal(t, true, exist)

	find, err := accountModel.FindOne(name)
	assert.NoError(t, err)
	assert.Equal(t, model.Uid, find.Uid)
}
