package service

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zty-f/Mini-Tiktok/controller"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ResourceBase 如果映射的域名和改了，需要更改这个配置
const ResourceBase = "http://192.168.0.101:8080/static/"

type PublishService struct {
}

// NewPublishServiceInstance 返回一个发布视频服务类的指针变量，可以方便调用该结构体的方法
func NewPublishServiceInstance() *PublishService {
	return &PublishService{}
}

// DoPublishVideo 发布视频
func (p *PublishService) DoPublishVideo(c *gin.Context, title string, userId int64) error {
	data, err := c.FormFile("data")
	if err != nil {
		return err
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		return err
	}
	playURL, vErr := joinResourceURL(ResourceBase, finalName)
	if vErr != nil {
		fmt.Printf("wrong join URL")
		fmt.Printf("Wrong URL is: %s", playURL)
	}
	// 获取视频第一帧作为封面图片
	reader := ExampleReadFrameAsJpeg("./public/"+finalName, 1)
	img, err1 := imaging.Decode(reader)
	if err1 != nil {
		return err1
	}
	err2 := imaging.Save(img, "./public/"+finalName+".jpeg")
	if err2 != nil {
		return err2
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
	err3 := videoDaoInstance.CreateVideoRecord(userId, playURL, coverURL, title)
	if err3 != nil {
		return err3
	}
	return nil
}

// DoPublishList 发布视频列表
func (p *PublishService) DoPublishList(userId, loginUserId int64) ([]controller.VideoVo, error) {
	videos, err := videoDaoInstance.QueryByOwner(userId)
	if err != nil {
		return nil, err
	}
	user, err1 := userDaoInstance.QueryUserById(userId)
	if err1 != nil {
		return nil, err1
	}
	isFollow, err3 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, userId)
	if err3 != nil {
		return nil, err3
	}
	curUser := &controller.UserVo{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	var PublishedList = make([]controller.VideoVo, len(videos))
	for i, _ := range PublishedList {
		var isFavorite bool
		actionType, err2 := favoriteDaoInstance.QueryActionTypeByUserIdAndVideoId(loginUserId, videos[i].Id)
		if err2 != nil {
			return nil, err2
		}
		if actionType == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		PublishedList[i] = controller.VideoVo{
			Id:            videos[i].Id,
			Author:        *curUser,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         videos[i].Title,
		}
	}
	return PublishedList, nil
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
