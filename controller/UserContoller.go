package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-pro/entity"
	"soft-pro/resp"
	"soft-pro/service"
)

type UserResponse struct {
	resp.Response
	User entity.User `json:"user"`
}

func Login(c *gin.Context) {
	phone := c.PostForm("telephone")
	password := c.PostForm("password")
	u, err := service.CheckLoginUser(phone, password)
	// 登陆失败
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkWithData(u, c)
}

func Register(c *gin.Context) {
	//获取参数
	u := entity.User{
		UserName: c.PostForm("user_name"),
		Phone:    c.PostForm("telephone"),
		Password: c.PostForm("password"),
		Role:     "user",
	}
	// 校验数据
	err := service.CheckRegisterUser(u)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	// 密码加密
	if service.InsertUser(u) != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	c.Redirect(http.StatusPermanentRedirect, "/login")
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("查询ID为:", id)
	u := service.GetUserByID(id)
	if u.ID == 0 {
		resp.FailWithMessage(resp.NotFindMsg, c)
	} else {
		resp.OkWithData(u, c)
	}
}
