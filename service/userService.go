package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zty-f/Mini-Tiktok/common"
	"github.com/zty-f/Mini-Tiktok/repository"
	"github.com/zty-f/Mini-Tiktok/utils"
	"strings"
)

type UserService struct {
}

// NewUserServiceInstance 返回一个用户服务类的指针变量，可以方便调用该结构体的方法
func NewUserServiceInstance() *UserService {
	return &UserService{}
}

// DoRegister 注册
func (u *UserService) DoRegister(userName, password string) (int64, string, error) {
	fmt.Printf("用户正在注册：" + userName + ":" + password)
	uErr := checkUserName(userName)
	pErr := checkPassword(password)
	if uErr != nil {
		return 0, "", uErr
	}
	if pErr != nil {
		return 0, "", pErr
	}
	flag, err1 := userDaoInstance.QueryIsContainsUserName(userName)
	if err1 != nil {
		return 0, "", err1
	}
	if flag {
		return 0, "", errors.New("用户名已存在，请创建一个独一无二的name吧！")
	}
	password = utils.MD5(password)
	//调用Dao层
	user, err := userDaoInstance.CreateByNameAndPassword(userName, password)
	if err != nil {
		return 0, "", err
	}
	// 因为用户名和密码拼接构成token也有可能构成重复，所以使用uuid生成唯一token
	token := uuid.NewV4()
	return user.Id, token.String(), nil
}

// DoLogin 登录
func (u *UserService) DoLogin(userName, password string) (*repository.User, string, error) {
	flag, err1 := userDaoInstance.QueryIsContainsUserName(userName)
	if err1 != nil {
		return nil, "", err1
	}
	if !flag {
		fmt.Println("用户名重复！")
		return nil, "", errors.New("用户名或者密码错误！")
	}
	user, err := userDaoInstance.QueryLoginInfo(userName)
	fmt.Println(user.Id)
	fmt.Println(user.Password)
	if err != nil {
		return nil, "", err
	}
	if !strings.EqualFold(utils.MD5(password), user.Password) {
		fmt.Println("密码错误！")
		return nil, "", errors.New("用户名或者密码错误！")
	}
	fmt.Printf("用户正在登录：" + userName + ":" + password)
	// 因为用户名和密码拼接构成token也有可能构成重复，所以使用uuid生成唯一token
	token := uuid.NewV4()
	return user, token.String(), nil
}

// DoUserInfo 获取用户信息
func (u *UserService) DoUserInfo(userId, loginUserId int64) (*common.UserVo, error) {
	user, err := userDaoInstance.QueryUserById(userId)
	if err != nil {
		return nil, err
	}
	favoriteCount, err1 := favoriteDaoInstance.QueryFavoriteCountByUserId(userId)
	if err1 != nil {
		return nil, err1
	}
	var totalFavorited = int64(0)
	if count, err4 := videoDaoInstance.QueryPublishCountByUserId(userId); err4 != nil {
		return nil, err4
	} else if count > 0 {
		totalFavorited, err = videoDaoInstance.QueryTotalFavoriteCountByUserId(userId)
	}
	if err != nil {
		return nil, err
	}
	isFollow, err3 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, userId)
	if err3 != nil {
		return nil, err3
	}
	loginUser := &common.UserVo{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          user.Avatar,
		Signature:       user.Signature,
		BackgroundImage: user.BackgroundImage,
		TotalFavorited:  totalFavorited,
		FavoriteCount:   favoriteCount,
	}
	return loginUser, nil
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
