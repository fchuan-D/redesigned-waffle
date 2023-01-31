package main

import (
	"github.com/gin-gonic/gin"
	"soft-pro/controller"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/")
	// basic apis
	apiRouter.GET("/get", controller.GetUser)
	// extra apis - I

	// extra apis - II

}
