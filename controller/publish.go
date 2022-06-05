package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PublishListResp struct {
	Response
	VideoList []VideoVo `json:"video_list,omitempty"`
}

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
	userId := OnlineUser[token].Id
	err := publishService.DoPublishVideo(c, title, userId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			1,
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "视频上传成功！！！！",
	})
	return
}

// PublishList 发布视频列表
func PublishList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	fmt.Println("查询发布视频列表用户id：" + c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			1,
			"用户id转换错误！",
		})
		return
	}
	loginUserId := OnlineUser[token].Id
	PublishedList, err1 := publishService.DoPublishList(userId, loginUserId)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			1,
			"服务端错误！",
		})
		return
	}
	c.JSON(http.StatusOK, PublishListResp{
		Response:  Response{StatusCode: 0, StatusMsg: "获取当前用户发布的视频列表成功！"},
		VideoList: PublishedList,
	})
	return
}
