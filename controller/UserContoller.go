package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"soft-pro/dao"
	"soft-pro/entity"
)

type UserResponse struct {
	entity.Response
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

func GetUser(c *gin.Context) {
	var u entity.User
	_ = dao.Db.AutoMigrate(entity.User{})
	dao.Db.First(&u)
	log.Println("==========>", u.UserName)
	c.JSON(http.StatusOK, UserResponse{
		UserName: u.UserName,
		Role:     u.Role,
	})
}
