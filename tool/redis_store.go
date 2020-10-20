package tool

import "github.com/go-redis/redis"

type RedisStore struct {
	client *redis.Client
}

var GlobalRedis RedisStore

func InitRedis() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})
	GlobalRedis = RedisStore{client: client}
	return &GlobalRedis
}
