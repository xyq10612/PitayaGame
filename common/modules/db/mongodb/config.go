package mongodb

import (
	"common/modules/db"
	"fmt"
)

type MongoConfig struct {
	db.Config
	MaxPoolSize int // 连接池大小
}

func (c *MongoConfig) GetConnURI() string {
	return fmt.Sprintf("mongodb://%s:%d", c.Host, c.Port)
}

func NewDefaultMongoConfig() *MongoConfig {
	return &MongoConfig{
		Config: db.Config{
			Host: "localhost",
			Port: 27017,
		},
		MaxPoolSize: 10,
	}
}
