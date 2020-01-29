package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis_main.host"),
		Password: viper.GetString("redis_main.password"),
		DB:       viper.GetInt("redis_main.db"),
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}
