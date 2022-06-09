package db2gorm

import (
	"github.com/zty-f/Mini-Tiktok/utils/db2gorm/gen"
)

func generateOne() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=true&loc=Local"

	//生成指定单表
	tblName := "Users"

	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}

func generateAll() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=true&loc=Local"

	gen.GenerateAll(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	})
}
