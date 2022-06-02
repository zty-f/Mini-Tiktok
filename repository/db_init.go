package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return err

}
