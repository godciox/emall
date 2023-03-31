package redisInit

import "github.com/go-redis/redis/v7"

var RDB *redis.Client

func RedisInit() {
	//初始化redis，连接地址和端口，密码，数据库名称
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
