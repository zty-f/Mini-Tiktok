package favorites

type Favorites struct {
	ID         int64
	UserID     int64
	VideoID    int64
	IsFavorite int8 `gorm:"default:0"`
}
