package accountModel

import (
	"common/constants"
	"common/helper"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type AccountModel struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Uid      string `bson:"uid"`
}

func (a *AccountModel) GetDBName() string {
	return constants.MongoAccountDB
}

func (a *AccountModel) GetCollectionName() string {
	return constants.MongoAccountCollection
}

func (a *AccountModel) New() error {
	return helper.GetMongo().InsertOne(a)
}

func Exist(name string) bool {
	return helper.GetMongo().Exist(constants.MongoAccountDB, constants.MongoAccountCollection, bson.M{"name": name})
}

func FindOne(name string) (AccountModel, error) {
	var model AccountModel
	err := helper.GetAccountDB().Collection(constants.MongoAccountCollection).FindOne(context.TODO(), bson.M{"name": name}).Decode(&model)
	return model, err
}
