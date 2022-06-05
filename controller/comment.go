package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentResponse struct {
	Response
	Comment CommentVo `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []CommentVo `json:"comment_list,omitempty"`
}

// CommentAction 评论功能
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	userId := OnlineUser[token].Id
	fmt.Printf("评论userId：%d==videoId：%d==actionType:%d\n", userId, videoId, actionType)
	if actionType == 1 {
		//新增评论
		commentText := c.Query("comment_text")
		//调用service层
		commentVo, err := commentService.DoAddCommentAction(userId, videoId, commentText)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{0, "新增评论成功！"},
			Comment:  *commentVo,
		})
	} else {
		//删除评论
		commentId, err4 := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err4 != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "服务端错误，评论操作失败！",
			})
			return
		}
		//调用service层
		err5 := commentService.DoDelCommentAction(videoId, commentId)
		if err5 != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "服务端错误，删除评论失败！",
			})
			return
		}
		c.JSON(http.StatusOK, Response{0, "删除评论成功！"})
	}
	return
}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	token := c.Query("token")
	loginUserId := OnlineUser[token].Id
	//调用service层
	commentList, err := commentService.DoCommentList(loginUserId, videoId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{0, "获取视频评论列表成功！"},
		CommentList: commentList,
	})
	return
}
