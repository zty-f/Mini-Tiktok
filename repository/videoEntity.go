package repository

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
	var videos []Video
	db.Order("create_time desc").Where("user_id=?", ownerId).Find(&videos)
	return videos
}

func (d *VideoDao) CreateVideoRecord(userId int64, playURL string, coverURL string, title string) error {
	var video = &Video{
		PlayUrl:  playURL,
		CoverUrl: coverURL,
		Title:    title,
		UserId:   userId,
	}
	db.Create(video)
	return nil
}

func (d *VideoDao) QueryFeedFlow(latestTime int64) []Video {
	var videos []Video
	db.Order("create_time desc").Limit(MaxListLength).Find(&videos)
	return videos
}
