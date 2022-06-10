package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var db *gorm.DB

// Init 初始化数据库连接操作
func Init() error {
	var err error
	dsn := "root:zty123456.@tcp(47.108.239.8:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return err

}
