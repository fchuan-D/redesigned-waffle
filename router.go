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

	userRouter := r.Group("/user")
	// extra apis - I
	userRouter.GET("/get/:id", jwt.JWT(), controller.GetUser)

	adminRouter := r.Group("/admin")
	// extra apis - II
	adminRouter.GET("/get/:id", jwt.JWT(), controller.GetUser)

}
