package dao

import (
	"log"
	"soft-pro/conf"
	"soft-pro/entity"
	"testing"
	"time"
)

// 数据库连接测试
func TestDao(t *testing.T) {
	conf.InitConfig("../")
	Init()

	var u entity.User
	Db.First(&u)
	log.Println("==========>", u.UserName)

	Rd.Set("test", "Hello World", 30*time.Second)
	s, _ := Rd.Get("test").Result()
	log.Println("==========>", s)
}
