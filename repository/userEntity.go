package repository

import "fmt"

type User struct {
	ID            int64  `gorm:"primary_key"`
	Name          string `gorm:"column:name;size:32;not null"`
	Password      string `gorm:"column:password;size:32;"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	IsFollow      bool   `json:"column:is_follow"`
}

type UserDao struct {
}

func NewUserDaoInstance() *UserDao {
	return &UserDao{}
}

func (u *UserDao) QueryUserById(userId int64) *User {
	var user = &User{}
	db.First(user, userId)
	return user
}

func (u *UserDao) CreateByNameAndPassword(name, password string) *User {
	var user = &User{
		ID:            0,
		Name:          name,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	db.Create(user)
	fmt.Println(user.ID)
	return user
}

func (u *UserDao) QueryLoginInfo(name, password string) *User {
	var user = &User{}
	db.Where("name = ? and password = ?", name, password).Find(user)
	return user
}
