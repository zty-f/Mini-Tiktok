package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Favorite struct {
	ID         int64
	UserID     int64
	VideoID    int64
	IsFavorite int32 `gorm:"default:0"`
}

type FavoriteDao struct {
}

func NewFavoriteDaoInstance() *FavoriteDao {
	return &FavoriteDao{}
}

func (f *FavoriteDao) ActionOfLike(userId int64, videoId int64, actionType int32) error {
	video := &Video{}
	favorite := &Favorite{
		UserID:     userId,
		VideoID:    videoId,
		IsFavorite: actionType,
	}
	// 点赞列表新增数据或修改标志位
	if err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(favorite).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("数据库不包含这个点赞数据，需要创建·····")
		//添加数据
		db.Select("user_id", "video_id", "is_favorite").Create(favorite)
	} else if err != nil {
		return err
	} else {
		fmt.Println("更新点赞标志位······")
		//更新标志位
		db.Model(favorite).Update("is_favorite", actionType)
	}
	if actionType == 1 {
		// 点赞视频的点赞数+1
		if err := db.Model(video).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
			return err
		}
	} else {
		if err := db.Model(video).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error; err != nil {
			return err
		}
	}
	return nil
}

func (f *FavoriteDao) QueryVideosIdByUserId(userId int64) []int64 {
	var ids []int64
	fmt.Println("通过userId查询点赞视频列表的videoId")
	db.Table("favorites").Select("video_id").Where("user_id = ?", userId).Find(&ids)
	return ids
}

func (f *FavoriteDao) QueryActionTypeByUserIdAndVideoId(userId, videoId int64) int32 {
	var actionType int32
	fmt.Println("通过userId+videoId查询点赞状态")
	db.Table("favorites").Select("is_favorite").Where("user_id = ? and video_id = ?", userId, videoId).Find(&actionType)
	return actionType
}
