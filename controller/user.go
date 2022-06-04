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

var OnlineUser = map[string]*UserVo{}

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

// Register 注册
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

// Login 登录
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
		IsFollow:      false,
	}

	//加入到online表里
	OnlineUser[tokenSb.String()] = loginUser
}

// UserInfo 获取用户详细信息
func UserInfo(c *gin.Context) {
	qid := c.Query("user_id")
	userId, err := strconv.ParseInt(qid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误！",
		})
		return
	}
	var userEntity = userDaoInstance.QueryUserById(userId)
	favoriteCount := favoriteDao.QueryFavoriteCountByUserId(userId)
	totalFavorited := videoDaoInstance.QueryTotalFavoriteCountByUserId(userId)
	loginUser := &UserVo{
		Id:              userEntity.Id,
		Name:            userEntity.Name,
		FollowCount:     userEntity.FollowCount,
		FollowerCount:   userEntity.FollowerCount,
		IsFollow:        false,
		Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
		Signature:       "冲冲冲，就快要做完了！",
		BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
		TotalFavorited:  totalFavorited,
		FavoriteCount:   favoriteCount,
	}
	c.JSON(http.StatusOK, UserResp{
		Response: Response{0, "获取登录用户详细信息成功！"},
		User:     *loginUser,
	})

	return
}

//检查用户名
func checkUserName(userName string) error {
	if len(userName) > MaxUsernameLen {
		return errors.New("username is too long")
	}

	return nil
}

//检查密码
func checkPassword(passWord string) error {
	if len(passWord) > MaxPasswordLen {
		return errors.New("password is too long")
	}
	return nil
}
