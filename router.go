package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/config"
	"github.com/zty-f/Mini-Tiktok/controller"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	// 不用拦截的接口
	r.GET("/douyin/feed/", controller.Feed)
	r.POST("/douyin/user/register/", controller.Register)
	r.POST("/douyin/user/login/", controller.Login)

	apiRouter := r.Group("/douyin")
	//配置拦截器
	apiRouter.Use(config.CheckToken)

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.PublishVideo)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)
}
