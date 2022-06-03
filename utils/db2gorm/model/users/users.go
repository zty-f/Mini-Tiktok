package users

type Users struct {
	ID            int64
	Name          string
	Password      string
	FollowCount   uint64 `gorm:"default:0"`
	FollowerCount uint64 `gorm:"default:0"`
	IsFollow      uint8  `gorm:"default:0"`
}
