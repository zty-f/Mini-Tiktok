package controller

import "github.com/zty-f/Mini-Tiktok/service"

var userService = service.NewUserServiceInstance()
var feedService = service.NewFeedServiceInstance()
var publishService = service.NewPublishServiceInstance()
var favoriteService = service.NewFavoriteServiceInstance()
