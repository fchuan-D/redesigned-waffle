package test

import (
	"log"
	"soft-pro/conf"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/middleware/redis"
	"testing"
	"time"
)

// 数据库连接测试
func TestDao(t *testing.T) {
	conf.InitConfig(".")
	dao.Init()

	var u entity.User
	dao.Db.First(&u)
	log.Println("==========>", u.UserName)

	redis.GetClient().Set("test", "Hello World", 30*time.Second)
	s, _ := redis.GetClient().Get("test").Result()
	log.Println("==========>", s)
}
