package service

import "github.com/zty-f/Mini-Tiktok/repository"

const MaxUsernameLen = 32
const MaxPasswordLen = 32

var userDaoInstance = repository.NewUserDaoInstance()
var videoDaoInstance = repository.NewVideoDaoInstance()
var favoriteDaoInstance = repository.NewFavoriteDaoInstance()
var relationDaoInstance = repository.NewRelationDaoInstance()
var commentDaoInstance = repository.NewCommentDaoInstance()
