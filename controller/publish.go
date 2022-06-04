package controller

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zty-f/Mini-Tiktok/repository"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PublishListResp struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
}

var videoDao = repository.NewVideoDaoInstance()

// ResourceBase 如果映射的域名和改了，需要更改这个配置
const ResourceBase = "http://192.168.0.101:8080/static/"

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	token := c.PostForm("token")
	if _, exists := OnlineUser[token]; !exists {
		fmt.Println("用户未登录········token:" + token)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "请先登录再进行后续操作，谢谢！",
		})
	}
	title := c.PostForm("title")
	fmt.Println(token + title)
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			1,
			err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	user := OnlineUser[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	playURL, vErr := joinResourceURL(ResourceBase, finalName)
	if vErr != nil {
		fmt.Printf("wrong join URL")
		fmt.Printf("Wrong URL is: %s", playURL)
	}
	// 获取视频第一帧作为封面图片
	reader := ExampleReadFrameAsJpeg("./public/"+finalName, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "封面图片解析错误",
		})
		return
	}
	err = imaging.Save(img, "./public/"+finalName+".jpeg")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "封面图片保存错误",
		})
		return
	}
	coverName := finalName + ".jpeg"
	coverURL, cErr := joinResourceURL(ResourceBase, coverName)
	if cErr != nil {
		fmt.Printf("wrong join URL")
		fmt.Printf("Wrong URL is: %s", coverURL)
	}
	fmt.Println("视频上传成功！")
	fmt.Println("playUrl:" + playURL)
	fmt.Println("coverURL:" + coverURL)
	videoDao.CreateVideoRecord(user.Id, playURL, coverURL, title)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "视频上传成功！！！！",
	})
	return
}

// 拼接字符串
func joinResourceURL(baseDomain, resourse string) (string, error) {
	var sb strings.Builder
	_, err := fmt.Fprintf(&sb, "%s/%s", baseDomain, resourse)
	if err != nil {
		fmt.Printf("joinResource fail %v", err)
		return "", err
	}
	return sb.String(), nil
}

// PublishList 发布视频列表
func PublishList(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	fmt.Println("查询发布视频列表用户id：" + c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			1,
			"用户id转换错误！",
		})
		return
	}
	var videos = videoDao.QueryByOwner(uid)
	var user = userDaoInstance.QueryUserById(uid)
	loginUser := &UserVo{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	var PublishedList = make([]VideoVo, len(videos))
	for i, _ := range PublishedList {
		var isFavorite bool
		actionType := favoriteDao.QueryActionTypeByUserIdAndVideoId(uid, videos[i].Id)
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		PublishedList[i] = VideoVo{
			Id:            videos[i].Id,
			Author:        *loginUser,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         videos[i].Title,
		}
	}
	c.JSON(http.StatusOK, PublishListResp{
		Response:  Response{StatusCode: 0, StatusMsg: "获取当前用户发布的视频列表成功！"},
		VideoList: PublishedList,
	})
	return

}

// ExampleReadFrameAsJpeg 视频解析成流
func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
