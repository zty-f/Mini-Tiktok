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
		favoriteCount := favoriteDao.QueryFavoriteCountByUserId(videoList[i].UserId)
		totalFavorited := videoDaoInstance.QueryTotalFavoriteCountByUserId(videoList[i].UserId)
		tmpUser := &UserVo{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        false,
			Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
			Signature:       "冲冲冲，就快要做完了！",
			BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
			TotalFavorited:  totalFavorited,
			FavoriteCount:   favoriteCount,
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
