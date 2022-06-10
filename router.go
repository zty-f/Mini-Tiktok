package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/config"
	"github.com/zty-f/Mini-Tiktok/controller"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	//不用拦截的接口组
	apiRouter1 := r.Group("/douyin")

	apiRouter1.GET("/feed/", controller.Feed)
	apiRouter1.POST("/user/register/", controller.Register)
	apiRouter1.POST("/user/login/", controller.Login)
	apiRouter1.POST("/publish/action/", controller.PublishVideo)
	apiRouter1.GET("/comment/list/", controller.CommentList)
	apiRouter1.GET("/favorite/list/", controller.FavoriteList)
	apiRouter1.GET("/publish/list/", controller.PublishList)

	//需要拦截的接口组
	apiRouter2 := r.Group("/douyin")
	//配置拦截器
	apiRouter2.Use(config.CheckToken)

	apiRouter2.GET("/user/", controller.UserInfo)
	apiRouter2.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter2.POST("/comment/action/", controller.CommentAction)
	apiRouter2.POST("/relation/action/", controller.RelationAction)
	apiRouter2.GET("/relation/follow/list/", controller.RelationFollowList)
	apiRouter2.GET("/relation/follower/list/", controller.RelationFollowerList)
}
