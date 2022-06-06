package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/controller"
	"github.com/zty-f/Mini-Tiktok/repository"
)

func main() {
	r := gin.Default()

	initRouter(r)
	if err := repository.Init(); err != nil {
		panic(err)
	}
	if err := controller.RedisInit(); err != nil {
		fmt.Printf("redis连接失败! err : %v\n", err)
		return
	}
	fmt.Println("redis连接成功！")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
