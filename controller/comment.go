package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
)

var commentDao = repository.NewCommentDaoInstance()

type CommentResponse struct {
	Response
	Comment CommentVo `json:"comment"`
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
		cid, err3 := commentDao.CreateComment(userId, videoId, commentText)
		comment, err4 := commentDao.QueryCommentById(cid)
		if err3 != nil || err4 != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "服务端错误，新增评论失败！",
			})
			return
		}
		// 按照2006-01-02 15:04:05这个固定来格式化
		createDate := comment.CreateTime.Format("01-02")
		fmt.Println(createDate)
		user := userDaoInstance.QueryUserById(userId)
		userVo := &UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		commentVo := &CommentVo{
			Id:         comment.ID,
			User:       *userVo,
			Content:    comment.Content,
			CreateDate: createDate,
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{0, "新增评论成功！"},
			Comment:  *commentVo,
		})
	} else {
		//删除评论
		commentId, err4 := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		err5 := commentDao.DeleteComment(commentId, videoId)
		if err4 != nil || err5 != nil {
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
