package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxysfive/Mini-Tiktok/controller"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action", controller.PublishVideo)
}
