package redis

import (
	"github.com/go-redis/redis/v7"
	"soft-pro/conf"
	"time"
)

func GetClient() (clint *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.GetConfig().RedisUrI,
		Password: conf.GetConfig().RedisPass,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic("连接Redis失败")
	}
	return client
}

/*
基于 go-redis
github地址:https://github.com/go-redis/redis
设置小时 demo
*/
func SetHour(key string, value string, i int64) (result *redis.StatusCmd) {
	client := GetClient()
	// 注意 time.Hour 要乘以一个具体的数值
	set := client.Set(key, value, time.Duration(i)*time.Hour)
	return set
}

/*
基于 go-redis
github地址:https://github.com/go-redis/redis
设置分钟 demo
*/
func SetMini(key string, value string, i int64) (result *redis.StatusCmd) {
	client := GetClient()
	// 注意 time.Minute 要乘以一个具体的数值
	set := client.Set(key, value, time.Duration(i)*time.Minute)
	return set
}

/*
基于 go-redis
github地址:https://github.com/go-redis/redis
get的 demo
*/
func Get(key string) string {
	c := GetClient()
	get := c.Get(key)
	result, _ := get.Result()
	return result
}
