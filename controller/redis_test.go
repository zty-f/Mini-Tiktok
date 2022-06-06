package controller

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	err := RedisInit()
	if err != nil {
		return
	}
	Rd.Set(Ctx, "test", "123", time.Second*5) //过期时间

}
