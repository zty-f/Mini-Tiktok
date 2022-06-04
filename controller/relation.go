package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
)

var relationDao = repository.NewRelationDaoInstance()

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, err1 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	userId := OnlineUser[token].Id
	if actionType == 1 {
		// 关注
		err := relationDao.CreateRelation(userId, toUserId)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "服务端错误，关注失败！",
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "关注成功！",
		})
	} else {
		// 取消关注
		err := relationDao.DeleteRelation(userId, toUserId)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "服务端错误，取消关注失败！",
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "取消关注成功！",
		})
	}
	return
}
