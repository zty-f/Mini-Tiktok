package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/repository"
	"net/http"
	"strconv"
)

var relationDao = repository.NewRelationDaoInstance()

type RelationListResponse struct {
	Response
	UserList []UserVo `json:"user_list,omitempty"`
}

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

func RelationFollowList(c *gin.Context) {
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "服务端错误，评论操作失败！",
		})
		return
	}
	// 根据用户id查询该用户关注的所有用户的id
	ids := relationDao.QueryFollowIdsByUserId(userId)
	if len(ids) == 0 {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{0, "还未关注任何用户，继续发现吧！"},
			UserList: nil,
		})
		return
	}
	// 根据用户的id查询用户信息
	users := userDaoInstance.QueryUsersByIds(ids)
	userList := make([]UserVo, len(users))
	for i, _ := range users {
		favoriteCount := favoriteDao.QueryFavoriteCountByUserId(users[i].Id)
		totalFavorited := videoDaoInstance.QueryTotalFavoriteCountByUserId(users[i].Id)
		isFollow := relationDao.QueryIsFollowByUserIdAndToUserId(userId, users[i].Id)
		userList[i] = UserVo{
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
	c.JSON(http.StatusOK, RelationListResponse{
		Response: Response{0, "获取关注列表成功！"},
		UserList: userList,
	})
	return
}
