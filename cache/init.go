package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/util/logger"
)

var RedisClient redis.UniversalClient

// todo：集群或单点连接配置化待完善
func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis_main.host"),
		Password:     viper.GetString("redis_main.password"),
		DB:           15,
		//DialTimeout:  500 * time.Microsecond, // 设置连接超时
		//ReadTimeout:  500 * time.Microsecond, // 设置读取超时
		//WriteTimeout: 500 * time.Microsecond, // 设置写入超时
	})

	_, err := client.Ping().Result()

	if err != nil {
		logger.Panic("conn_to_redis_cluster_failed", logger.Fields{"err": err})
	}

	RedisClient = client
}

func TestRedisConn() {
	result, err := RedisClient.Set("foo", "bar", 0).Result()
	if err != nil {
		fmt.Printf("error at set foo bar,err=%v\n", err)
		panic(err)
	}

	fmt.Println("result1:" + result)

	result2, err2 := RedisClient.Set("hello", "world", 0).Result()
	if err2 != nil {
		fmt.Printf("error at set hello world,err=%v\n", err)
		panic(err)
	}

	fmt.Println("result1:" + result2)
}
