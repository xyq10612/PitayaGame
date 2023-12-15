package helper

import (
	"common/constants"
	"common/modules/db/redis"
	"github.com/topfreegames/pitaya/v2"
)

var r *redis.RedisStorage

func GetRedis() *redis.RedisStorage {
	if r == nil {
		module, err := pitaya.DefaultApp.GetModule(constants.RedisModule)
		if err != nil {
			panic(err)
		}

		r = module.(*redis.RedisStorage)
	}

	return r
}
