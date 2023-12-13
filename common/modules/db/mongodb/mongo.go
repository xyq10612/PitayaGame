package mongodb

import (
	"context"
	"errors"
	"github.com/topfreegames/pitaya/v2/modules"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStorage struct {
	modules.Base
	*mongo.Client
	config MongoConfig
}

func NewMongoStorage(config MongoConfig) *MongoStorage {
	return &MongoStorage{
		config: config,
	}
}

func (m *MongoStorage) Init() error {
	m.Connect()
	return nil
}

func (m *MongoStorage) Connect() {
	opt := options.Client().ApplyURI(m.config.GetConnURI())
	if m.config.Username != "" && m.config.Password != "" {
		opt.SetAuth(options.Credential{
			Username: m.config.Username,
			Password: m.config.Password,
		})
	}

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		panic(err)
	}

	m.Client = client

	if err = m.TestPing(); err != nil {
		panic(err)
	}
}

func (m *MongoStorage) TestPing() error {
	return m.Ping(context.TODO(), readpref.Primary())
}

func (m *MongoStorage) Close() {
	_ = m.Disconnect(context.TODO())
}

func (m *MongoStorage) GetCollection(dbName, collection string) *mongo.Collection {
	return m.Database(dbName).Collection(collection)
}

func (m *MongoStorage) InsertOne(c Collection) error {
	_, err := m.GetCollection(c.GetDBName(), c.GetCollectionName()).InsertOne(context.TODO(), c)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStorage) Exist(dbName, collection string, filter interface{}) bool {
	opt := &options.FindOneOptions{Projection: filter}
	r := m.GetCollection(dbName, collection).FindOne(context.TODO(), filter, opt)
	return !errors.Is(mongo.ErrNoDocuments, r.Err())
}
