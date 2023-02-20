package controller

import (
	"github.com/gin-gonic/gin"
	"soft-pro/entity"
	"soft-pro/resp"
	"soft-pro/service"
)

// GET /order/info/:OrderID 获取订单信息
func OrderInfo(c *gin.Context) {
	oid := c.Param("OrderID")
	user, err := service.GetOrderByID(oid)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
	} else {
		resp.OkWithData(user, c)
	}
}

// GET /order/list 获取用户所有订单
func OrderList(c *gin.Context) {
	u, _ := c.Get("user")
	user, err := service.GetOrdersByID(u.(entity.User).ID)
	if err != nil {
		resp.FailWithMessage(resp.UserNotExistErrorMsg, c)
	} else {
		resp.OkWithData(user, c)
	}
}

// POST /order/create 创建订单
func CreateOrder(c *gin.Context) {
	//获取参数
	order := entity.Order{
		Status:  false,
		Amount:  c.PostForm("amount"),
		Charge:  c.PostForm("charge"),
		Type:    c.PostForm("type"),
		UserID:  c.PostForm("userID"),
		PointID: c.PostForm("pointID"),
	}
	err := service.CreateOrder(order)
	if err != nil {
		resp.FailWithMessage(resp.NotOkMsg, c)
	}
	resp.OkResult(c)
}

// GET /order/pay/:OrderID 支付订单
func PayOrder(c *gin.Context) {
	//获取参数
	oid := c.Param("OrderID")
	u, _ := c.Get("user")
	err := service.PayOrder(oid, u.(entity.User).ID)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
		return
	}
	resp.OkResult(c)
}
