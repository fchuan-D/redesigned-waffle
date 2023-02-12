package test

import (
	"log"
	"soft-pro/conf"
	"soft-pro/dao"
	"soft-pro/entity"
	"testing"
	"time"
)

// 数据库连接测试
func TestDao(t *testing.T) {
	conf.InitConfig("../")
	dao.Init()

	var u entity.User
	dao.Db.First(&u)
	log.Println("==========>", u.UserName)

	dao.Rd.Set("test", "Hello World", 30*time.Second)
	s, _ := dao.Rd.Get("test").Result()
	log.Println("==========>", s)
}
