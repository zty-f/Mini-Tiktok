package repository

type Comment struct {
	UserId      int64  `gorm:"primaryKey;autoIncrement:false"`
	VideoId     int64  `gorm:"primaryKey;autoIncrement:false"`
	CommentTest string `gorm:"size:255"`
	CreateTime  int64
}

type CommentDao struct {
}

// NewCommentDaoInstance 返回一个评论表实体类的指针变量，可以方便调用该结构体的方法
func NewCommentDaoInstance() *CommentDao {
	return &CommentDao{}
}

func (c *CommentDao) QueryCommentById(uid, vid int64) *Comment {
	//根据符合主键找到符合条件的comments
	var comment = &Comment{
		UserId:      uid,
		VideoId:     vid,
		CommentTest: "",
		CreateTime:  0,
	}
	db.Find(comment)
	return comment
}
