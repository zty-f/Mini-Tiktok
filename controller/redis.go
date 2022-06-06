package controller

//redis
import (
	"context"
	"github.com/go-redis/redis/v8"
)

// Rd 定义一个全局变量
var Rd *redis.Client
var Ctx = context.Background()

func RedisInit() (err error) {
	Rd = redis.NewClient(&redis.Options{
		Addr:     "47.108.239.8:6379", // 指定
		Password: "123456",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = Rd.Ping(Ctx).Result()
	return err
}
