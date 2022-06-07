package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestName(t *testing.T) {
	// NewV1 根据当前时间戳和 MAC 地址返回 UUID。
	fmt.Println(uuid.NewV1())
	// NewV2 基于 POSIX UID/GID 返回 DCE 安全 UUID。
	fmt.Println(uuid.NewV2(byte(1)))
	// NewV3 根据命名空间 UUID 和名称的 MD5 哈希返回 UUID。
	fmt.Println(uuid.NewV3(uuid.NamespaceDNS, "xxx"))
	// NewV4 返回随机生成的 UUID。
	fmt.Println(uuid.NewV4())
}
