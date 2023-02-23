package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-pro/entity"
	"soft-pro/middleware/jwt"
	"soft-pro/resp"
	"soft-pro/service"
	"strconv"
)

type LoginResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}

// POST /login 用户登录
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
	resp.OkWithData(LoginResponse{
		User:  u,
		Token: token,
	}, c)
}

// POST /register 用户注册
func Register(c *gin.Context) {
	//获取参数
	u := entity.User{
		UserName: c.PostForm("userName"),
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

// GET /user/info 获取用户信息
func UserInfo(c *gin.Context) {
	u, _ := c.Get("user")
	user, err := service.GetUserByID(u.(entity.User).ID)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
	} else {
		resp.OkWithData(user, c)
	}
}

// POST /user/update 更新用户信息
func UpdateUser(c *gin.Context) {
	var u entity.User
	err := c.ShouldBind(&u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 校验数据
	if err := service.CheckUpdateUser(u); err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	// 更新数据
	if err := service.InsertUser(u); err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkWithData(u, c)
}

// POST /user/update/bal 更新用户余额
func UpdateBal(c *gin.Context) {
	bal := c.PostForm("balance")
	changeBal, err := strconv.ParseFloat(bal, 64)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	u, _ := c.Get("user")
	err = service.UpdateBal(u.(entity.User).ID, changeBal)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkResult(c)
}
