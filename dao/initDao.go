package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"soft-pro/conf"
	"soft-pro/entity"
	"time"
)

var Db *gorm.DB

// 初始化数据库
func Init() {
	initMysql()
}

// 初始化Mysql数据库
func initMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)
	var err error
	dsn := conf.GetConfig().MysqlUrI
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	createTable()
}

// 映射结构体为数据库表
func createTable() {
	_ = Db.AutoMigrate(
		entity.User{},
		entity.ChargePoint{},
		entity.ChargeStation{},
		entity.Oreder{},
	)
}
