package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

var RedisClient redis.UniversalClient

// todo：集群或单点连接配置化待完善
func Redis() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: func() ([]redis.ClusterSlot, error) {
			return []redis.ClusterSlot{
				{
					Start: 0,
					End: 5460,
					Nodes: []redis.ClusterNode{
						{
							Addr: "192.168.127.130:7001",
						},
						{
							Addr: "192.168.127.130:7004",
						},
					},
				},
				{
					Start: 5461,
					End: 10922,
					Nodes: []redis.ClusterNode{
						{
							Addr: "192.168.127.130:7002",
						},
						{
							Addr: "192.168.127.130:7005",
						},
					},
				},
				{
					Start: 10923,
					End: 16383,
					Nodes: []redis.ClusterNode{
						{
							Addr: "192.168.127.130:7003",
						},
						{
							Addr: "192.168.127.130:7006",
						},
					},
				},
			}, nil
		},
		Password:     viper.GetString("redis_main.password"),
		DialTimeout:  500 * time.Microsecond, // 设置连接超时
		ReadTimeout:  500 * time.Microsecond, // 设置读取超时
		WriteTimeout: 500 * time.Microsecond, // 设置写入超时
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
