package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var OnlineUser = map[string]*UserVo{}

type RegisterResp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
type UserResp struct {
	Response
	User UserVo `json:"user"`
}

// Register 注册
func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	//调用Service层
	var userId, token, err = userService.DoRegister(userName, password)
	if err != nil {
		c.JSON(http.StatusOK, RegisterResp{
			Response: Response{StatusCode: 1,
				StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResp{
		Response: Response{0, "注册成功！"},
		UserId:   userId, //不知道该怎么写了
		Token:    token,
	})
	return
}

// Login 登录
func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	//调用service层
	user, token, err := userService.DoLogin(userName, password)
	if user.Id <= 0 || err != nil {
		c.JSON(http.StatusOK, RegisterResp{
			Response: Response{StatusCode: 1,
				StatusMsg: "用户名或密码错误！"},
		})
		return
	}
	var loginUser = &UserVo{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}

	//加入到online表里
	OnlineUser[token] = loginUser
	c.JSON(http.StatusOK, RegisterResp{
		Response: Response{0, "登录成功！"},
		UserId:   user.Id,
		Token:    token,
	})
	return
}

// UserInfo 获取用户详细信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	qid := c.Query("user_id")
	userId, err := strconv.ParseInt(qid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	loginUserId := OnlineUser[token].Id
	//调用service层
	userInfo, err1 := userService.DoUserInfo(userId, loginUserId)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	c.JSON(http.StatusOK, UserResp{
		Response: Response{0, "获取登录用户详细信息成功！"},
		User:     *userInfo,
	})
	return
}
