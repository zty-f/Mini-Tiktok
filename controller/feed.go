package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
	"time"
)

var videoDaoInstance = repository.NewVideoDaoInstance()

type FeedResponse struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
	NextTime  int64     `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	strTime := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(strTime, 10, 64)
	if err != nil {
		fmt.Printf("wrong parse string result is: %v", latestTime)
		latestTime = time.Now().Unix()
	}
	token := c.Query("token")
	if _, exists := onlineUser[token]; !exists {
		fmt.Println("用户未登录········")
	}
	videoList := videoDaoInstance.QueryFeedFlow(latestTime)
	videoListResp := make([]VideoVo, len(videoList))
	fmt.Println("获取视频流成功！")
	for i, _ := range videoList {
		user := userDaoInstance.QueryUserById(videoList[i].UserId)
		loginUser := &UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		videoListResp[i] = VideoVo{
			Id:            videoList[i].Id,
			Author:        *loginUser,
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			IsFavorite:    videoList[i].IsFavorite,
			Title:         videoList[i].Title,
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "获取视频流成功！"},
		VideoList: videoListResp,
		NextTime:  time.Now().Unix(),
	})
}
