package service

import (
	"errors"
	"github.com/zty-f/Mini-Tiktok/common"
)

type RelationService struct {
}

// NewRelationServiceInstance 返回一个关注服务类的指针变量，可以方便调用该结构体的方法
func NewRelationServiceInstance() *RelationService {
	return &RelationService{}
}

// DoRelationAction 关注
func (r *RelationService) DoRelationAction(userId, toUserId, actionType int64) error {
	if actionType == 1 {
		// 关注
		err := relationDaoInstance.CreateRelation(userId, toUserId)
		if err != nil {
			return err
		}
	} else {
		// 取消关注
		err := relationDaoInstance.DeleteRelation(userId, toUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

// DoRelationFollowList 获取关注列表
func (r *RelationService) DoRelationFollowList(userId, loginUserId int64) ([]common.UserVo, error) {
	// 根据用户id查询该用户关注的所有用户的id
	ids, err := relationDaoInstance.QueryFollowIdsByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("还未关注任何用户，继续发现吧！")
	}
	// 根据用户的id查询用户信息
	users, err4 := userDaoInstance.QueryUsersByIds(ids)
	if err4 != nil {
		return nil, err4
	}
	userList := make([]common.UserVo, len(users))
	for i, _ := range users {
		favoriteCount, err1 := favoriteDaoInstance.QueryFavoriteCountByUserId(users[i].Id)
		if err1 != nil {
			return nil, err1
		}
		var totalFavorited = int64(0)
		if count, err2 := videoDaoInstance.QueryPublishCountByUserId(users[i].Id); err2 != nil {
			return nil, err2
		} else if count > 0 {
			totalFavorited, err = videoDaoInstance.QueryTotalFavoriteCountByUserId(users[i].Id)
		}
		if err != nil {
			return nil, err
		}
		isFollow, err3 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, users[i].Id)
		if err3 != nil {
			return nil, err3
		}
		userList[i] = common.UserVo{
			Id:              users[i].Id,
			Name:            users[i].Name,
			FollowCount:     users[i].FollowCount,
			FollowerCount:   users[i].FollowerCount,
			Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
			Signature:       "冲冲冲，就快要做完了！",
			BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
			IsFollow:        isFollow,
			TotalFavorited:  totalFavorited,
			FavoriteCount:   favoriteCount,
		}
	}
	return userList, nil
}

// DoRelationFollowerList 获取粉丝列表
func (r *RelationService) DoRelationFollowerList(userId, loginUserId int64) ([]common.UserVo, error) {
	ids, err := relationDaoInstance.QueryFollowerIdsByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("还没有一个粉丝，有点可怜，继续创作吧！")
	}
	// 根据用户的id查询用户信息
	users, err1 := userDaoInstance.QueryUsersByIds(ids)
	if err1 != nil {
		return nil, err1
	}
	userList := make([]common.UserVo, len(users))
	for i, _ := range users {
		favoriteCount, err2 := favoriteDaoInstance.QueryFavoriteCountByUserId(users[i].Id)
		if err2 != nil {
			return nil, err2
		}
		var totalFavorited = int64(0)
		if count, err4 := videoDaoInstance.QueryPublishCountByUserId(users[i].Id); err4 != nil {
			return nil, err4
		} else if count > 0 {
			totalFavorited, err = videoDaoInstance.QueryTotalFavoriteCountByUserId(users[i].Id)
		}
		if err != nil {
			return nil, err
		}
		isFollow, err4 := relationDaoInstance.QueryIsFollowByUserIdAndToUserId(loginUserId, users[i].Id)
		if err4 != nil {
			return nil, err4
		}
		userList[i] = common.UserVo{
			Id:              users[i].Id,
			Name:            users[i].Name,
			FollowCount:     users[i].FollowCount,
			FollowerCount:   users[i].FollowerCount,
			Avatar:          "https://s3.bmp.ovh/imgs/2022/05/04/345d42da2a13020b.jpg",
			Signature:       "冲冲冲，就快要做完了！",
			BackgroundImage: "https://s3.bmp.ovh/imgs/2022/05/04/29ccf3f609f3e5f2.jpg",
			IsFollow:        isFollow,
			TotalFavorited:  totalFavorited,
			FavoriteCount:   favoriteCount,
		}
	}
	return userList, nil
}
