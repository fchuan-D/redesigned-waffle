package main

import (
	"github.com/gin-gonic/gin"
	"soft-pro/controller"
	"soft-pro/middleware"
	"soft-pro/middleware/jwt"
)

func initRouter(r *gin.Engine) {
	r.Use(middleware.Cors())

	basicRouter := r.Group("/")
	// basic apis
	basicRouter.POST("/register", controller.Register)
	basicRouter.POST("/login", controller.Login)
	basicRouter.POST("/station/list", controller.StationList)
	basicRouter.GET("/points/:StationID", controller.PointList)

	// extra apis - I
	userRouter := r.Group("/user")
	userRouter.Use(jwt.UserJWT())
	userRouter.GET("/info", controller.UserInfo)
	userRouter.POST("/update", controller.UpdateUser)
	userRouter.GET("/order/info/:OrderID", controller.OrderInfo)
	userRouter.GET("/order/list", controller.OrderList)
	userRouter.POST("/order/create", controller.CreateOrder)
	userRouter.GET("/order/pay/:OrderID", controller.PayOrder)

	// extra apis - II
	adminRouter := r.Group("/admin")
	adminRouter.Use(jwt.AdminJWT())
	adminRouter.GET("/info", controller.UserInfo)
}
