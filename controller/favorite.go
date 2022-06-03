package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
)

var favoriteDao = repository.NewFavoriteDaoInstance()

func Action(c *gin.Context) {
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	videoId, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	fmt.Printf("点赞userId：%d==videoId：%d==actionType:%d\n", userId, videoId, actionType)
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	if _, exists := onlineUser[token]; !exists {
		fmt.Println("用户未登录········")
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "请先登录再进行点赞！",
		})
		return
	}
	err := favoriteDao.ActionOfLike(userId, videoId, int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，点赞失败！",
		})
		return
	}
	c.JSON(http.StatusOK,
		Response{
			StatusCode: 0,
			StatusMsg:  "更新点赞状态成功！",
		})
	return
}
