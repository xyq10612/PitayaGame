package redis

import (
	"common/modules/db"
	"fmt"
)

type RedisConfig struct {
	db.Config
	Protocol int
}

func (rc *RedisConfig) GetConnURI() string {
	if rc.Protocol == 0 {
		rc.Protocol = 3
	}

	if rc.Password == "" {
		return fmt.Sprintf("redis://%s:%d?db=0&protocol=%d", rc.Host, rc.Port, rc.Protocol)
	}

	return fmt.Sprintf("redis://:%s@%s:%d?db=0&protocol=%d", rc.Password, rc.Host, rc.Port, rc.Protocol)
}

func NewDefaultRedisConfig() *RedisConfig {
	return &RedisConfig{
		Config: db.Config{
			Host: "localhost",
			Port: 6379,
		},
		Protocol: 3,
	}
}
