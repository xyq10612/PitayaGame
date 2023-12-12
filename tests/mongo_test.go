package tests

import (
	"common/modules/db"
	"common/modules/db/mongodb"
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type Account struct {
	Name string `bson:"name"`
	Pwd  string `bson:"pwd"`
	Uid  string `bson:"uid"`
}

func registerAccount(mongo *mongodb.MongoStorage, name, pwd, uid string) error {
	account := &Account{
		name, pwd, uid,
	}
	_, err := mongo.GetCollection("tests", "account").InsertOne(context.TODO(), account)
	return err
}

func queryAccount(mongo *mongodb.MongoStorage, name string) (Account, error) {
	var account Account
	err := mongo.GetCollection("tests", "account").FindOne(context.TODO(), bson.M{"name": name}).Decode(&account)
	return account, err
}

func TestRegister(t *testing.T) {
	config := &mongodb.MongoConfig{
		Config: db.Config{
			Host:     "localhost",
			Port:     27017,
			Username: "debugeve",
			Password: "develop2023",
		},
	}
	s := mongodb.NewMongoStorage(*config)
	s.Init()
	s.Connect()
	err := registerAccount(s, "test", "pwdtest", "test001uid")
	assert.NoError(t, err)
	account, err := queryAccount(s, "test")
	assert.NoError(t, err)
	assert.Equal(t, "test001uid", account.Uid)
}
