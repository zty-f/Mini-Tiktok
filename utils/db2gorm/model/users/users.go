package users

type Users struct {
	ID              int64
	Name            string
	Password        string
	FollowCount     uint64 `gorm:"default:0"`
	FollowerCount   uint64 `gorm:"default:0"`
	Avatar          string
	Signature       string
	BackgroundImage string
}
