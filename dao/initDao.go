package dao

import (
	"github.com/go-redis/redis/v7"
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
var Rd *redis.Client

// 初始化数据库
func Init() {
	initMysql()
	initRedis()
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
	_ = Db.AutoMigrate(entity.User{})
	_ = Db.AutoMigrate(entity.ChargePoint{})
}

// 初始化Redis数据库
func initRedis() {
	Rd = redis.NewClient(&redis.Options{
		Addr:     conf.GetConfig().RedisUrI,
		Password: conf.GetConfig().RedisPass,
	})
	_, err := Rd.Ping().Result()
	if err != nil {
		panic(err)
	}
}
