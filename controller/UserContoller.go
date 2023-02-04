package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-pro/entity"
	"soft-pro/service"
)

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("查询ID为:", id)
	u := service.Search(id)
	if u.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{
			StatusCode: http.StatusUnprocessableEntity,
			StatusMsg:  "查询不到该数据...",
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "查询成功",
			},
			User: u,
		})
	}

}
