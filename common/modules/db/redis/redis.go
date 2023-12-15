package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/topfreegames/pitaya/v2/modules"
)

type RedisStorage struct {
	modules.Base
	*redis.Client
	config RedisConfig
}

func NewRedisStorage(config RedisConfig) *RedisStorage {
	return &RedisStorage{
		config: config,
	}
}

func (r *RedisStorage) Init() error {
	r.Connect()
	return nil
}

func (r *RedisStorage) Connect() {
	uri := r.config.GetConnURI()
	opts, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	r.Client = redis.NewClient(opts)

	if err := r.TestPing(); err != nil {
		panic(err)
	}
}

func (r *RedisStorage) TestPing() error {
	_, err := r.Client.Ping(context.TODO()).Result()
	return err
}

func (r *RedisStorage) Close() {
	_ = r.Client.Close()
}
