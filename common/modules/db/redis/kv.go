package redis

type KV interface {
	GetRedisKey() string
}
