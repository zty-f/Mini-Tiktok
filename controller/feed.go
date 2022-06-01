package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mxysfive/Mini-Tiktok/repository"
	"net/http"
	"strconv"
	"time"
)

var videoDaoInstance = repository.NewVideoDaoInstance()

type FeedRequest struct {
	LatestTime int64  `json:"latest_time,omitempty"`
	Token      string `json:"token,omitempty"`
}

type VideoResp struct {
	Id            int64           `json:"id"`
	Author        repository.User `json:"author"`
	PlayUrl       string          `json:"play_url"`
	CoverUrl      string          `json:"cover_url"`
	FavoriteCount int64           `json:"favorite_count"`
	CommentCount  int64           `json:"comment_count"`
	IsFavorite    bool            `json:"is_favorite"`
	Title         string          `json:"title"`
}

type FeedResponse struct {
	Response
	VideoList []VideoResp `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
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
	videoListResp := make([]VideoResp, len(videoList))
	fmt.Println("获取视频流成功！")
	for i, _ := range videoList {
		user := userDaoInstance.QueryUserById(videoList[i].UserId)
		videoListResp[i] = VideoResp{
			Id:            videoList[i].Id,
			Author:        *user,
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

func PackVideoList() (videoList []repository.Video) {
	//查video的数据
	return nil
}
