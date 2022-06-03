package videos

type Videos struct {
	ID            int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64 `gorm:"default:0"`
	CommentCount  int64 `gorm:"default:0"`
	Title         string
	UserID        int64
	IsFavorite    int8      `gorm:"default:0"`
	CreateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
