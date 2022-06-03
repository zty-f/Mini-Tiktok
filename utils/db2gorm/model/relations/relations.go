package relations

type Relations struct {
	ID          int64
	UserID      int64 `gorm:"default:0"`
	FollowingID int64 `gorm:"default:0"`
	IsFollow    int8  `gorm:"default:0"`
}
