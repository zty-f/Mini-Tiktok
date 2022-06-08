package repository

import "errors"

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
func (u *UserDao) QueryUserById(userId int64) (*User, error) {
	var user = &User{}
	if err := db.First(user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateByNameAndPassword 通过用户名和密码创建新用户
func (u *UserDao) CreateByNameAndPassword(name, password string) (*User, error) {
	var user = &User{
		Id:            0,
		Name:          name,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	var count int64
	if err := db.Table("users").Where("name = ?", name).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户名已存在，请创建一个独一无二的name吧！")
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// QueryLoginInfo 通过用户名和密码查询是否包含此用户
func (u *UserDao) QueryLoginInfo(name string, password string) (*User, error) {
	var user = &User{}
	if err := db.Where("name = ? and password = ?", name, password).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// QueryUsersByIds 通过一组用户id查询一组用户信息
func (u *UserDao) QueryUsersByIds(ids []int64) ([]User, error) {
	var users []User
	if err := db.Find(&users, ids).Error; err != nil {
		return nil, err
	}
	return users, nil
}
