package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/common"
	"net/http"
	"strconv"
)

type FavoriteResponse struct {
	common.Response
	VideoList []common.VideoVo `json:"video_list,omitempty"`
}

// FavoriteAction 点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	userId := OnlineUser[token].Id
	//调用service层
	err := favoriteService.DoFavoriteAction(userId, videoId, int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	c.JSON(http.StatusOK,
		common.Response{
			StatusCode: 0,
			StatusMsg:  "更新点赞状态成功！",
		})
	return
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	var loginUserId int64
	if len(token) == 0 {
		loginUserId = 0
	} else {
		loginUserId = OnlineUser[token].Id
	}
	//调用service层
	videoListResp, err := favoriteService.DoFavoriteList(userId, loginUserId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, FavoriteResponse{
		Response:  common.Response{0, "获取点赞视频列表成功！"},
		VideoList: videoListResp,
	})
	return
}
