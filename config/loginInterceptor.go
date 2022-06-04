package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zty-f/Mini-Tiktok/controller"
	"net/http"
)

func CheckToken(c *gin.Context) {
	token := c.Query("token")
	if _, exists := controller.OnlineUser[token]; !exists {
		fmt.Println("用户未登录········")
		c.JSON(http.StatusOK, controller.Response{
			StatusCode: 1,
			StatusMsg:  "请先登录再进行后续操作，谢谢！",
		})
		c.Abort()
	} else {
		c.Next()
	}
}
