package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
)

func main() {
	r := gin.Default()

	initRouter(r)
	if err := repository.Init(); err != nil {
		panic(err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
