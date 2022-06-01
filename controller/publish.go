package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mxysfive/Mini-Tiktok/repository"
	"net/http"
	"path/filepath"
	"strings"
)

type PublishListResp struct {
	Response
	VideoList []repository.Video
}

var videoDao = repository.NewVideoDaoInstance()

// ResourceBase 如果映射的域名和改了，需要更改这个配置
const ResourceBase = "http://5kt3855788.zicp.vip/static/"

func PublishVideo(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	if _, exists := onlineUser[token]; !exists {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user is not exists",
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			1,
			err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	user := onlineUser[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "upload success",
		})
	}
	playURL, vErr := joinResourceURL(ResourceBase, finalName)
	if vErr != nil {
		fmt.Printf("wrong join URL")
		fmt.Printf("Wrong URL is: %s", playURL)
	}
	coverName := ""
	coverURL, cErr := joinResourceURL(ResourceBase, coverName)
	if cErr != nil {
		fmt.Printf("wrong join URL")
		fmt.Printf("Wrong URL is: %s", coverURL)
	}
	videoDao.CreateVideoRecord(user.Id, playURL, coverURL, title)
	return
}

func joinResourceURL(baseDomain, resourse string) (string, error) {
	var sb strings.Builder
	_, err := fmt.Fprintf(&sb, "%s/%s", baseDomain, resourse)
	if err != nil {
		fmt.Printf("joinResource fail %v", err)
		return "", err
	}
	return sb.String(), nil
}

func PublishList(c *gin.Context) {
	// uid := c.Query("user_id")
	utoken := c.Query("token")

	if _, exists := onlineUser[utoken]; !exists {
		c.JSON(http.StatusOK, Response{
			1,
			"user is not exists",
		})
		return
	}
	user := onlineUser[utoken]
	var videos = videoDao.QueryByOwner(user.Id)
	var PublishedList = make([]repository.Video, len(videos))
	for i, _ := range PublishedList {
		PublishedList[i] = repository.Video{
			Id:            videos[i].Id,
			UserID:        videos[i].UserID,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    false, //自己给自己都是false吧我猜的
			Title:         videos[i].Title,
		}
	}
	c.JSON(http.StatusOK, PublishedList)
	return

}
