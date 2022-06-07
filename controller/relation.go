package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/common"
	"net/http"
	"strconv"
)

type RelationListResponse struct {
	common.Response
	UserList []common.UserVo `json:"user_list,omitempty"`
}

// RelationAction 关注
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, err1 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	userId := OnlineUser[token].Id
	// 调用service层
	err := relationService.DoRelationAction(userId, toUserId, actionType)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，修改关注状态失败！",
		})
		return
	}
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "修改关注状态成功！",
	})
	return
}

// RelationFollowList 获取关注列表
func RelationFollowList(c *gin.Context) {
	token := c.Query("token")
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	loginUserId := OnlineUser[token].Id
	// 调用service层
	userList, err := relationService.DoRelationFollowList(userId, loginUserId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, RelationListResponse{
		Response: common.Response{0, "获取关注列表成功！"},
		UserList: userList,
	})
	return
}

// RelationFollowerList 获取粉丝列表
func RelationFollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	loginUserId := OnlineUser[token].Id
	// 调用service层
	userList, err := relationService.DoRelationFollowerList(userId, loginUserId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, RelationListResponse{
		Response: common.Response{0, "获取粉丝列表成功！"},
		UserList: userList,
	})
	return
}
