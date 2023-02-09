package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soft-pro/entity"
	"soft-pro/resp"
	"soft-pro/service"
)

type UserResponse struct {
	resp.Response
	User entity.User `json:"user"`
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("查询ID为:", id)
	u := service.Search(id)
	if u.ID == 0 {
		resp.FailWithMessage(resp.NotFindMsg, c)
	} else {
		resp.OkWithData(u, c)
	}
}
