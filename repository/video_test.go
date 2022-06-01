package repository

import (
	"fmt"
	"testing"
	"time"
)

var videoDao = NewVideoDaoInstance()

func TestVideoDao_QueryFeedFlow(t *testing.T) {
	if err := Init(); err != nil {
		panic("connect fail")
	}
	db.AutoMigrate(&Video{})
	var videos = videoDao.QueryFeedFlow(time.Now().Unix())
	for _, v := range videos {
		fmt.Println(v)
	}
	if err := db.Model(videos).Association("User").Error; err != nil {
		fmt.Println("------------------------------")
	}

	for _, v := range videos {
		fmt.Println(v)
	}
}
