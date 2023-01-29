package dao

import (
	"log"
	"soft-pro/entity"
	"testing"
)

func TestDao(t *testing.T) {
	Init()
	var u entity.User
	Db.AutoMigrate(entity.User{})
	Db.First(&u)
	log.Println("==========>", u.UserName)
}
