package service

import (
	"errors"
	"fmt"
	"github.com/zty-f/Mini-Tiktok/common"
)

type FavoriteService struct {
}

// NewFavoriteServiceInstance 返回一个点赞服务类的指针变量，可以方便调用该结构体的方法
func NewFavoriteServiceInstance() *FavoriteService {
	return &FavoriteService{}
}

// DoFavoriteAction 点赞
func (f *FavoriteService) DoFavoriteAction(userId, videoId int64, actionType int32) error {
	fmt.Printf("点赞userId：%d==videoId：%d==actionType:%d\n", userId, videoId, actionType)
	flag, err := favoriteDaoInstance.QueryFavoriteByUserIdAndVideoId(userId, videoId)
	if err != nil {
		return err
	}
	if flag {
		err = favoriteDaoInstance.UpdateFavorite(userId, videoId, actionType)
		if err != nil {
			return err
		}
	} else {
		err = favoriteDaoInstance.CreateFavorite(userId, videoId, actionType)
		if err != nil {
			return err
		}
	}
	return nil
}

// DoFavoriteList 获取点赞列表
func (f *FavoriteService) DoFavoriteList(userId, loginUserId int64) ([]common.VideoVo, error) {
	fmt.Printf("获取点赞视频列表userId：%d\n", userId)
	ids, err := favoriteDaoInstance.QueryVideosIdByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("该用户未点赞任何视频！")
	}
	videoList, err1 := videoDaoInstance.QueryByIds(ids)
	if err1 != nil {
		return nil, err1
	}
	videoListResp := make([]common.VideoVo, len(videoList))
	fmt.Println("获取点赞视频列表成功！")
	for i, _ := range videoList {
		var isFavorite bool
		user, err2 := userDaoInstance.QueryUserById(videoList[i].UserId)
		if err2 != nil {
			return nil, err2
		}
		actionType, err3 := favoriteDaoInstance.QueryActionTypeByUserIdAndVideoId(loginUserId, videoList[i].Id)
		if err3 != nil {
			return nil, err3
		}
		isFollow, err4 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, userId)
		if err4 != nil {
			return nil, err4
		}
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		tmpUser := &common.UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
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
	return videoListResp, nil
}
