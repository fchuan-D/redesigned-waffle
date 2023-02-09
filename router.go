package main

import (
	"github.com/gin-gonic/gin"
	"soft-pro/controller"
	"soft-pro/middleware"
)

func initRouter(r *gin.Engine) {
	r.Use(middleware.Cors())

	apiRouter := r.Group("/")
	// basic apis
	apiRouter.GET("/get/:id", controller.GetUser)
	// extra apis - I

	// extra apis - II

}
