package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"soft-pro/dao"
	"soft-pro/entity"
)

func GetUser(c *gin.Context) {
	var u entity.User
	_ = dao.Db.AutoMigrate(entity.User{})
	dao.Db.First(&u)
	log.Println("==========>", u.UserName)
}
