package repository

type Comment struct {
	UserId      int64  `gorm:"primaryKey;autoIncrement:false"`
	VideoId     int64  `gorm:"primaryKey;autoIncrement:false"`
	CommentTest string `gorm:"size:255"`
	CreateTime  int64
}

type CommentDao struct {
}

func NewCommentDaoInstance() *CommentDao {
	return &CommentDao{}
}

func (c *CommentDao) QueryCommentByPKs(uid, vid int64) *Comment {
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
