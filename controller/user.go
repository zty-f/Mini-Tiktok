package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
	"strings"
)

var onlineUser = map[string]*UserVo{}

const MaxUsernameLen = 32
const MaxPasswordLen = 32

type RegisterResp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
type UserResp struct {
	Response
	User UserVo `json:"user"`
}

var userDaoInstance = repository.NewUserDaoInstance()

func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	uErr := checkUserName(userName)
	pErr := checkPassword(password)
	fmt.Printf("用户正在注册：" + userName + ":" + password)
	if uErr != nil {
		c.JSON(http.StatusOK, RegisterResp{
			Response: Response{StatusCode: 1,
				StatusMsg: uErr.Error()},
		})
		return
	}
	if pErr != nil {
		c.JSON(http.StatusOK, RegisterResp{
			Response: Response{StatusCode: 1,
				StatusMsg: pErr.Error()},
		})
		return
	}
	//调用Dao层
	var user = userDaoInstance.CreateByNameAndPassword(userName, password)

	var tokenSb strings.Builder
	fmt.Fprintf(&tokenSb, "%s%s", userName, password)
	c.JSON(http.StatusOK, RegisterResp{
		Response: Response{0, "注册成功！"},
		UserId:   user.Id, //不知道该怎么写了
		Token:    tokenSb.String(),
	})
	return
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	var user = userDaoInstance.QueryLoginInfo(userName, password)
	if user.Id <= 0 {
		c.JSON(http.StatusOK, RegisterResp{
			Response: Response{StatusCode: 1,
				StatusMsg: "用户名或密码错误！"},
		})
		return
	}
	fmt.Printf("用户正在登录：" + userName + ":" + password)
	var tokenSb strings.Builder
	fmt.Fprintf(&tokenSb, "%s%s", userName, password)
	c.JSON(http.StatusOK, RegisterResp{
		Response: Response{0, "登录成功！"},
		UserId:   user.Id,
		Token:    tokenSb.String(),
	})

	var loginUser = &UserVo{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}

	//加入到online表里
	onlineUser[tokenSb.String()] = loginUser
}

func UserInfo(c *gin.Context) {
	qid := c.Query("user_id")
	utoken := c.Query("token") //判断用户是否登录
	if _, exists := onlineUser[utoken]; !exists {
		c.JSON(http.StatusOK, Response{
			StatusCode: 4001,
			StatusMsg:  "用户未登录请重新登陆！",
		})
		return
	}

	userId, err := strconv.Atoi(qid)
	if err != nil {
		fmt.Printf("Function of atoi in UserInfo fail %v", err)
	}
	var userEntity = userDaoInstance.QueryUserById(int64(userId))
	fmt.Println("登录用户: ", userEntity)
	loginUser := &UserVo{
		Id:            userEntity.Id,
		Name:          userEntity.Name,
		FollowCount:   userEntity.FollowCount,
		FollowerCount: userEntity.FollowerCount,
		IsFollow:      userEntity.IsFollow,
	}
	c.JSON(http.StatusOK, UserResp{
		Response: Response{0, "获取登录用户详细信息成功！"},
		User:     *loginUser,
	})

	return
}

func checkUserName(userName string) error {
	if len(userName) > MaxUsernameLen {
		return errors.New("username is too long")
	}

	return nil
}

func checkPassword(passWord string) error {
	if len(passWord) > MaxPasswordLen {
		return errors.New("password is too long")
	}
	return nil
}
