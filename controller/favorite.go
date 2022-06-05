package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
)

var favoriteDao = repository.NewFavoriteDaoInstance()

type FavoriteResponse struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
}

// FavoriteAction 点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	userId := OnlineUser[token].Id
	fmt.Printf("点赞userId：%d==videoId：%d==actionType:%d\n", userId, videoId, actionType)
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

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	fmt.Printf("获取点赞视频列表userId：%d\n", userId)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	ids := favoriteDao.QueryVideosIdByUserId(userId)
	if len(ids) == 0 {
		fmt.Println("该用户未点赞任何视频！")
		c.JSON(http.StatusOK, FavoriteResponse{
			Response:  Response{0, "该用户未点赞任何视频！"},
			VideoList: nil,
		})
	}
	loginUserId := OnlineUser[token].Id
	videoList := videoDaoInstance.QueryByIds(ids)
	videoListResp := make([]VideoVo, len(videoList))
	fmt.Println("获取点赞视频列表成功！")
	for i, _ := range videoList {
		var isFavorite bool
		user := userDaoInstance.QueryUserById(videoList[i].UserId)
		actionType := favoriteDao.QueryActionTypeByUserIdAndVideoId(loginUserId, videoList[i].Id)
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		loginUser := &UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		videoListResp[i] = VideoVo{
			Id:            videoList[i].Id,
			Author:        *loginUser,
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         videoList[i].Title,
		}
	}
	c.JSON(http.StatusOK, FavoriteResponse{
		Response:  Response{0, "获取点赞视频列表成功！"},
		VideoList: videoListResp,
	})
}
