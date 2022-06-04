package comments

import "time"

type Comments struct {
	ID         int64
	UserID     string
	VideoID    string
	Content    string
	CreateDate string
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
