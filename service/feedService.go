package service

import (
	"fmt"
	"github.com/zty-f/Mini-Tiktok/common"
)

type FeedService struct {
}

// NewFeedServiceInstance 返回一个视频流服务类的指针变量，可以方便调用该结构体的方法
func NewFeedServiceInstance() *FeedService {
	return &FeedService{}
}

// DoFeed 获取视频流
func (f *FeedService) DoFeed(loginUserId int64, latestTime string) ([]common.VideoVo, error) {
	videoList, err := videoDaoInstance.QueryFeedFlow(latestTime)
	if err != nil {
		return nil, err
	}
	videoListResp := make([]common.VideoVo, len(videoList))
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
		var totalFavorited = int64(0)
		if count, err4 := videoDaoInstance.QueryPublishCountByUserId(videoList[i].UserId); err4 != nil {
			return nil, err4
		} else if count > 0 {
			totalFavorited, err = videoDaoInstance.QueryTotalFavoriteCountByUserId(videoList[i].UserId)
		}
		if err != nil {
			return nil, err
		}
		isFollow, err5 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, videoList[i].UserId)
		if err5 != nil {
			return nil, err5
		}
		tmpUser := &common.UserVo{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
			TotalFavorited:  totalFavorited,
			FavoriteCount:   favoriteCount,
		}
		videoListResp[i] = common.VideoVo{
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
