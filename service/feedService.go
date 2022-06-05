package service

import (
	"fmt"
	"github.com/zty-f/Mini-Tiktok/controller"
)

type FeedService struct {
}

// NewFeedServiceInstance 返回一个视频流服务类的指针变量，可以方便调用该结构体的方法
func NewFeedServiceInstance() *FeedService {
	return &FeedService{}
}

// DoFeed 获取视频流
func (f *FeedService) DoFeed(loginUserId int64, latestTime string) ([]controller.VideoVo, error) {
	videoList, err := videoDaoInstance.QueryFeedFlow(latestTime)
	if err != nil {
		return nil, err
	}
	videoListResp := make([]controller.VideoVo, len(videoList))
	fmt.Println("获取视频流成功！")
	for i, _ := range videoList {
		var isFavorite bool
		user, err1 := userDaoInstance.QueryUserById(videoList[i].UserId)
		if err1 != nil {
			return nil, err1
		}
		actionType, err2 := favoriteDaoInstance.QueryActionTypeByUserIdAndVideoId(loginUserId, videoList[i].Id)
		if err2 != nil {
			return nil, err2
		}
		fmt.Println(actionType)
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		favoriteCount, err3 := favoriteDaoInstance.QueryFavoriteCountByUserId(videoList[i].UserId)
		if err3 != nil {
			return nil, err3
		}
		totalFavorited, err4 := videoDaoInstance.QueryTotalFavoriteCountByUserId(videoList[i].UserId)
		if err4 != nil {
			return nil, err4
		}
		isFollow, err5 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, videoList[i].UserId)
		if err5 != nil {
			return nil, err5
		}
		tmpUser := &controller.UserVo{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
			Signature:       "冲冲冲，就快要做完了！",
			BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
			TotalFavorited:  totalFavorited,
			FavoriteCount:   favoriteCount,
		}
		videoListResp[i] = controller.VideoVo{
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
	return videoListResp, err
}
