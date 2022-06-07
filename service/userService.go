package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zty-f/Mini-Tiktok/common"
	"github.com/zty-f/Mini-Tiktok/repository"
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
	user, err := userDaoInstance.QueryLoginInfo(userName, password)
	if err != nil {
		return nil, "", err
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
	totalFavorited, err2 := videoDaoInstance.QueryTotalFavoriteCountByUserId(userId)
	if err2 != nil {
		return nil, err2
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
		Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
		Signature:       "冲冲冲，就快要做完了！",
		BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
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
