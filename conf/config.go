package conf

import "time"

// 数据库连接字段
// 想要正确的处理 time.Time,需要带上 parseTime 参数，
// 要支持完整的 UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4
var Dsn = "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

// 设置 Redis数据热度消散时间。
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24
var ExpireTime = time.Hour * 48
