package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-pro/entity"
	"soft-pro/middleware/jwt"
	"soft-pro/resp"
	"soft-pro/service"
)

type UserResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}

func Login(c *gin.Context) {
	phone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 校验登陆信息
	u, err := service.CheckLoginUser(phone, password)
	// 登陆失败
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}

	// 登陆成功后签发 token
	token, err := jwt.GenerateToken(u.ID, u.UserName)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	// 缓存至 Redis
	service.SaveBufferToRd(token, u)
	resp.OkWithData(UserResponse{
		User:  u,
		Token: token,
	}, c)
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
	// 新增用户
	if service.InsertUser(u) != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	c.Redirect(http.StatusPermanentRedirect, "/login")
}

func GetUser(c *gin.Context) {
	u := c.MustGet("user").(entity.User)
	fmt.Println("当前登录用户ID:", u.ID)
	if u.ID == 0 {
		resp.FailWithMessage(resp.NotFindMsg, c)
	} else {
		resp.OkWithData(u, c)
	}
}
