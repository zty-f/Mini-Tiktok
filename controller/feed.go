package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"time"
)

var videoDaoInstance = repository.NewVideoDaoInstance()

type FeedResponse struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
	NextTime  int64     `json:"next_time,omitempty"`
}

// Feed 获取视频流
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	if latestTime == "" {
		fmt.Printf("wrong parse string result is: %v", latestTime)
		latestTime = time.Now().String()
	}
	var loginUser UserVo
	token := c.Query("token")
	if _, exists := OnlineUser[token]; !exists {
		fmt.Println("用户未登录········")
	} else {
		loginUser = *OnlineUser[token]
	}
	// 调用service层
	videoListResp, err := feedService.DoFeed(loginUser.Id, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "获取视频流成功！"},
		VideoList: videoListResp,
		NextTime:  time.Now().Unix(),
	})
	return
}
