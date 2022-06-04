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

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	fmt.Printf("获取评论列表列表的videoId：%d\n", videoId)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	comments, err2 := commentDao.QueryCommentsByVideoId(videoId)
	if err2 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	if len(comments) == 0 {
		fmt.Println("该视频还没有用户评论，欢迎评论^_^！")
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{0, "该视频还没有用户评论，欢迎评论^_^！"},
			CommentList: nil,
		})
		return
	}
	commentList := make([]CommentVo, len(comments))
	fmt.Println("获取该视频评论列表成功！")
	for i, _ := range comments {
		// 按照2006-01-02 15:04:05这个固定来格式化
		createDate := comments[i].CreateTime.Format("01-02")
		user := userDaoInstance.QueryUserById(comments[i].UserID)
		userVo := &UserVo{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		commentList[i] = CommentVo{
			Id:         comments[i].ID,
			User:       *userVo,
			Content:    comments[i].Content,
			CreateDate: createDate,
		}
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{0, "获取视频评论列表成功！"},
		CommentList: commentList,
	})
	return
}
