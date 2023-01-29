package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"soft-pro/dao"
)

func main() {
	// 环境初始化
	InitDeps()

	// Gin
	r := gin.Default()
	initRouter(r)

	// pprof
	pprof.Register(r)

	// localhost:8080
	r.Run()
}

// 加载项目环境
func InitDeps() {
	// 初始化数据库
	dao.Init()
}