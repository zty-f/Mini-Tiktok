package repository

import (
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

// NewFavoriteDaoInstance 返回一个点赞表实体类的指针变量，可以方便调用该结构体的方法
func NewFavoriteDaoInstance() *FavoriteDao {
	return &FavoriteDao{}
}

// QueryFavoriteByUserIdAndVideoId 查询是否包含对应点赞关联记录
func (f *FavoriteDao) QueryFavoriteByUserIdAndVideoId(userId int64, videoId int64) (bool, error) {
	var count int64
	if err := db.Table("favorites").Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateFavorite 创建点赞关联记录
func (f *FavoriteDao) CreateFavorite(userId int64, videoId int64, actionType int32) error {
	video := &Video{}
	favorite := &Favorite{
		UserID:     userId,
		VideoID:    videoId,
		IsFavorite: actionType,
	}
	tx := db.Begin()
	//添加数据
	if err := tx.Select("user_id", "video_id", "is_favorite").Create(favorite).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 点赞视频的点赞数+1
	if err := tx.Model(video).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateFavorite 更新点赞关联记录
func (f *FavoriteDao) UpdateFavorite(userId int64, videoId int64, actionType int32) error {
	video := &Video{}
	favorite := &Favorite{
		UserID:     userId,
		VideoID:    videoId,
		IsFavorite: actionType,
	}
	fmt.Println("更新点赞标志位······")
	tx := db.Begin()
	//更新标志位
	if err := tx.Model(favorite).Where("user_id = ? and video_id = ?", userId, videoId).Update("is_favorite", actionType).Error; err != nil {
		tx.Rollback()
		return err
	}
	if actionType == 1 {
		// 点赞视频的点赞数+1
		if err := tx.Model(video).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(video).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// QueryVideosIdByUserId 通过用户id查询查询该用户点赞的所有视频对应的视频id列表
func (f *FavoriteDao) QueryVideosIdByUserId(userId int64) ([]int64, error) {
	var ids []int64
	fmt.Println("通过userId查询点赞视频列表的videoId")
	if err := db.Table("favorites").Select("video_id").Where("user_id = ? and is_favorite = ?", userId, 1).Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

// QueryActionTypeByUserIdAndVideoId 通过用户id和视频id获取该用户对于这个视频是否点赞的状态码
func (f *FavoriteDao) QueryActionTypeByUserIdAndVideoId(userId, videoId int64) (int32, error) {
	var actionType int32
	fmt.Println("通过userId+videoId查询点赞状态")
	if err := db.Table("favorites").Select("is_favorite").Where("user_id = ? and video_id = ?", userId, videoId).Find(&actionType).Error; err != nil {
		return 0, err
	}
	return actionType, nil
}

// QueryFavoriteCountByUserId 根据用户id查询用户点赞视频的数量
func (f *FavoriteDao) QueryFavoriteCountByUserId(userId int64) (int64, error) {
	var favoriteCount int64
	fmt.Println("通过userId查询点赞视频列表的videoId")
	if err := db.Model(&Favorite{}).Where("user_id = ? and is_favorite = ?", userId, 1).Count(&favoriteCount).Error; err != nil {
		return 0, err
	}
	return favoriteCount, nil
}
