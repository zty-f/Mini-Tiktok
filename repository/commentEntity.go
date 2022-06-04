package repository

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         int64
	UserID     int64
	VideoID    int64
	Content    string
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CommentDao struct {
}

// NewCommentDaoInstance 返回一个评论表实体类的指针变量，可以方便调用该结构体的方法
func NewCommentDaoInstance() *CommentDao {
	return &CommentDao{}
}

// CreateComment 新增评论返回评论id
func (c *CommentDao) CreateComment(userId, videoId int64, commentText string) (int64, error) {
	video := &Video{}
	comment := &Comment{
		UserID:  userId,
		VideoID: videoId,
		Content: commentText,
	}
	if err := db.Select("user_id", "video_id", "content").Create(&comment).Error; err != nil {
		return comment.ID, err
	}
	// 对应视频评论数+1
	if err := db.Model(video).Where("id = ? ", videoId).Update("comment_count", gorm.Expr("comment_count+ ?", 1)).Error; err != nil {
		return comment.ID, err
	}
	// 返回comment对象只包含传入字段以及主键id
	return comment.ID, nil
}

// DeleteComment 删除评论
func (c *CommentDao) DeleteComment(commentId, videoId int64) error {
	video := &Video{}
	comment := &Comment{}
	err := db.Delete(comment, commentId).Error
	if err != nil {
		return err
	}
	// 对应视频评论数-1
	if err1 := db.Model(video).Where("id = ? ", videoId).Update("comment_count", gorm.Expr("comment_count- ?", 1)).Error; err1 != nil {
		return err1
	}
	return nil
}

// QueryCommentById 通过id查询评论
func (c *CommentDao) QueryCommentById(commentId int64) (*Comment, error) {
	comment := &Comment{}
	err := db.First(comment, commentId).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// QueryCommentsByVideoId 通过视频id查询该视频所有评论
func (c *CommentDao) QueryCommentsByVideoId(videoId int64) ([]Comment, error) {
	var comments []Comment
	fmt.Println("通过videoId查询所有评论")
	err := db.Order("create_time desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}
