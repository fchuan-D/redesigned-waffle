package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"soft-pro/conf"
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
	err := r.Run(conf.GetConfig().Port)
	if err != nil {
		return
	}
}

// 加载项目环境
func InitDeps() {
	// 初始化配置
	conf.InitConfig(".")
	// 初始化数据库
	dao.Init()
}
