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

	// extra apis - I
	userRouter := r.Group("/user")
	userRouter.Use(jwt.UserJWT())
	userRouter.GET("/get/:id", controller.GetUser)

	// extra apis - II
	adminRouter := r.Group("/admin")
	adminRouter.Use(jwt.AdminJWT())
	adminRouter.GET("/get/:id", controller.GetUser)

}
