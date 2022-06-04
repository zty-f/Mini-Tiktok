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
}

type VideoDao struct {
}

// NewVideoDaoInstance 返回一个视频实体类的指针变量，可以方便调用该结构体的方法
func NewVideoDaoInstance() *VideoDao {
	return &VideoDao{}
}

// QueryByOwner 通过用户id查询该用户发布的所有视频
func (d *VideoDao) QueryByOwner(ownerId int64) []Video {
	var videos []Video
	db.Order("create_time desc").Where("user_id=?", ownerId).Find(&videos)
	return videos
}

// QueryByIds 通过一组视频id获取对应的视频列表
func (d *VideoDao) QueryByIds(ids []int64) []Video {
	var videos []Video
	db.Find(&videos, ids)
	return videos
}

// CreateVideoRecord 通过传入参数创建新的视频记录
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

// QueryFeedFlow 通过当前时间查询这个时间前新发布的50条视频，逆序输出
func (d *VideoDao) QueryFeedFlow(latestTime int64) []Video {
	var videos []Video
	db.Order("create_time desc").Limit(MaxListLength).Find(&videos)
	return videos
}
