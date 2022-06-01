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

type FeedResponse struct {
	Response
	VideoList []repository.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
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
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user is not exists",
		})
		return
	} else {
		//token 所对应的用户存在
		videoList := videoDaoInstance.QueryFeedFlow(latestTime)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videoList,
			NextTime:  time.Now().Unix(),
		})
	}

}

func PackVideoList() (videoList []repository.Video) {
	//查video的数据
	return nil
}
