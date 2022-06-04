package repository

import "fmt"

type User struct {
	Id            int64  `gorm:"primary_key"`
	Name          string `gorm:"column:name;size:32;not null"`
	Password      string `gorm:"column:password;size:32;"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

type UserDao struct {
}

// NewUserDaoInstance 返回一个用户实体类的指针变量，可以方便调用该结构体的方法
func NewUserDaoInstance() *UserDao {
	return &UserDao{}
}

// QueryUserById 通过用户id查询详细的用户信息
func (u *UserDao) QueryUserById(userId int64) *User {
	var user = &User{}
	db.First(user, userId)
	return user
}

// CreateByNameAndPassword 通过用户名和密码创建新用户
func (u *UserDao) CreateByNameAndPassword(name, password string) *User {
	var user = &User{
		Id:            0,
		Name:          name,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	db.Create(user)
	fmt.Println(user.Id)
	return user
}

// QueryLoginInfo 通过用户名和密码查询是否包含此用户
func (u *UserDao) QueryLoginInfo(name string, password string) *User {
	var user = &User{}
	db.Where("name = ? and password = ?", name, password).Find(user)
	return user
}
