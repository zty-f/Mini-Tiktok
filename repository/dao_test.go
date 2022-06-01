package repository

import (
	"fmt"
	"testing"
)

func TestDAO(t *testing.T) {
	err := Init()
	if err != nil {
		fmt.Println(err.Error())
		panic("fail in connect with database")
	}
	db.AutoMigrate(&User{}, &Video{})
	var uList = []User{

		{
			//ID:            0,
			Name:          "mxy",
			Password:      "123456",
			FollowCount:   0,
			FollowerCount: 0,
		},
		{
			//ID:            0,
			Name:          "msy",
			Password:      "msy123456",
			FollowCount:   0,
			FollowerCount: 1,
		},
	}

	var video = &Video{
		ID:            0,
		PlayUrl:       "www.woshi.com",
		CoverUrl:      "1",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "test",
		UserID:        1,
	}

	db.Create(&uList[0])
	fmt.Println(uList[0].ID)
	db.Create(&uList[1])
	fmt.Println(uList[1].ID)

	db.Create(video)
	fmt.Println(video.ID)

}

func TestBelongTo(t *testing.T) {
	err := Init()
	if err != nil {
		fmt.Println(err.Error())
		panic("fail in connect with database")
	}
	var videos []Video
	var user User
	db.Model(&user).Where("id = ?", "1").Find(&user)

	db.Model(&videos).Where("user_id = ?", "1").Association("User")
	for _, v := range videos {
		fmt.Println(v)
	}

	//db.Model(video).Association("User")

	return
}

func TestAutoMigrate(t *testing.T) {
	err := Init()
	if err != nil {
		panic(err)
	}

}
