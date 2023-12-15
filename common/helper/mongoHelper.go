package helper

import (
	"common/constants"
	"common/modules/db/mongodb"
	"github.com/topfreegames/pitaya/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var m *mongodb.MongoStorage

func GetMongo() *mongodb.MongoStorage {
	if m == nil {
		module, err := pitaya.DefaultApp.GetModule(constants.MongoDBModule)
		if err != nil {
			panic(err)
		}

		m = module.(*mongodb.MongoStorage)
	}

	return m
}

func GetGameDB() *mongo.Database {
	return GetMongo().Database(constants.MongoGameDB)
}

func GetAccountDB() *mongo.Database {
	return GetMongo().Database(constants.MongoAccountDB)
}
