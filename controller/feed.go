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

// Feed 获取视频流
func Feed(c *gin.Context) {
	strTime := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(strTime, 10, 64)
	if err != nil {
		fmt.Printf("wrong parse string result is: %v", latestTime)
		latestTime = time.Now().Unix()
	}
	var loginUser UserVo
	token := c.Query("token")
	if _, exists := OnlineUser[token]; !exists {
		fmt.Println("用户未登录········")
	} else {
		loginUser = *OnlineUser[token]
	}
	videoList := videoDaoInstance.QueryFeedFlow(latestTime)
	videoListResp := make([]VideoVo, len(videoList))
	fmt.Println("获取视频流成功！")
	for i, _ := range videoList {
		var isFavorite bool
		user := userDaoInstance.QueryUserById(videoList[i].UserId)
		actionType := favoriteDao.QueryActionTypeByUserIdAndVideoId(loginUser.Id, videoList[i].Id)
		fmt.Println(actionType)
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		tmpUser := &UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		videoListResp[i] = VideoVo{
			Id:            videoList[i].Id,
			Author:        *tmpUser,
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         videoList[i].Title,
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "获取视频流成功！"},
		VideoList: videoListResp,
		NextTime:  time.Now().Unix(),
	})
}
