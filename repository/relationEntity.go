package repository

import "gorm.io/gorm"

type Relation struct {
	ID          int64
	UserID      int64 `gorm:"default:0"`
	FollowingID int64 `gorm:"default:0"`
}

type RelationDao struct {
}

// NewRelationDaoInstance 返回一个评论表实体类的指针变量，可以方便调用该结构体的方法
func NewRelationDaoInstance() *RelationDao {
	return &RelationDao{}
}

// CreateRelation 新增关注记录信息
func (r *RelationDao) CreateRelation(userId, toUserId int64) error {
	user := &User{}
	relation := &Relation{
		UserID:      userId,
		FollowingID: toUserId,
	}
	if err := db.Select("user_id", "following_id").Create(&relation).Error; err != nil {
		return err
	}
	// 对应用户粉丝数+1
	if err := db.Model(user).Where("id = ? ", toUserId).Update("follower_count", gorm.Expr("follower_count+ ?", 1)).Error; err != nil {
		return err
	}
	// 对应用户关注数量+1
	if err := db.Model(user).Where("id = ? ", userId).Update("follow_count", gorm.Expr("follow_count+ ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRelation 删除关注记录信息
func (r *RelationDao) DeleteRelation(userId, toUserId int64) error {
	user := &User{}
	relation := &Relation{
		UserID:      userId,
		FollowingID: toUserId,
	}
	if err := db.Where("user_id = ? and following_id = ?", userId, toUserId).Delete(relation).Error; err != nil {
		return err
	}
	// 对应用户粉丝数-1
	if err := db.Model(user).Where("id = ? ", toUserId).Update("follower_count", gorm.Expr("follower_count- ?", 1)).Error; err != nil {
		return err
	}
	// 对应用户关注数量-1
	if err := db.Model(user).Where("id = ? ", userId).Update("follow_count", gorm.Expr("follow_count- ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
