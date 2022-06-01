package repository

import (
	"fmt"
	"testing"
)

var usrDao = NewUserDaoInstance()

func TestUserDao_CreateByNameAndPassword(t *testing.T) {
	if err := Init(); err != nil {
		panic(err)
	}
	var result = usrDao.CreateByNameAndPassword("name", "password")
	fmt.Println(result.ID)
}

func TestUserDao_QueryUserById(t *testing.T) {
	if err := Init(); err != nil {
		panic(err)
	}
	var result = usrDao.QueryUserById(3)
	fmt.Println(result.ID)
}

func TestUserDao_QueryLoginInfo(t *testing.T) {
	if err := Init(); err != nil {
		panic(err)
	}
	usrName := "mxy"
	password := "123456"
	var result = usrDao.QueryLoginInfo(usrName, password)
	fmt.Println(result.ID)
}
