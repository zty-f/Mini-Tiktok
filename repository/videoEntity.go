package repository

import (
	"errors"
)

const MaxListLength = 30

type Video struct {
	Id            int64  `gorm:"primaryKey"`
	PlayUrl       string `gorm:"size:64"`
	CoverUrl      string `gorm:"size:64"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title;size:32"`
	UserId        int64  `gorm:"column:user_id"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
}

type VideoDao struct {
}

func NewVideoDaoInstance() *VideoDao {
	return &VideoDao{}
}

func (d *VideoDao) QueryByOwner(ownerId int64) []Video {
	//在用户查看自己的发布视频时使用，feed接口不用这个
	var video = &Video{
		Id:            0,
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "",
		UserId:        ownerId,
		IsFavorite:    false,
	}
	var videos []Video
	db.Where(video).Find(videos)
	return videos
}

func (d *VideoDao) CreateVideoRecord(userId int64, playURL string, coverURL string, title string) error {
	var video = &Video{
		Id:            0,
		PlayUrl:       playURL,
		CoverUrl:      coverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		UserId:        userId,
		IsFavorite:    false,
	}
	if video.Id == 0 {
		return errors.New("failure in create video record")
	}
	return nil

}

func (d *VideoDao) QueryFeedFlow(latestTime int64) []Video {
	var videos []Video
	db.Order("create_time desc").Limit(MaxListLength).Find(&videos)
	return videos
}
