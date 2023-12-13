package helper

import (
	"common/constants"
	"common/modules/db/mongodb"
	"github.com/topfreegames/pitaya/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongodb.MongoStorage

func GetMongo() *mongodb.MongoStorage {
	if db == nil {
		m, err := pitaya.DefaultApp.GetModule(constants.MongoDBModule)
		if err != nil {
			panic(err)
		}

		db = m.(*mongodb.MongoStorage)
	}

	return db
}

func GetGameDB() *mongo.Database {
	return GetMongo().Database(constants.MongoGameDB)
}

func GetAccountDB() *mongo.Database {
	return GetMongo().Database(constants.MongoAccountDB)
}
