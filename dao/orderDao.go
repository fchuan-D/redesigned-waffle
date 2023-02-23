package dao

import (
	"fmt"
	"soft-pro/entity"
	"time"
)

// 根据 OrderID获取订单数据
func GetOrderByID(id any) entity.Order {
	var o entity.Order
	Db.Find(&o, id)
	return o
}

// 获取当前用户所有订单
func GetOrdersByUserID(UserID any) []entity.Order {
	var os []entity.Order
	Db.Find(&os, "user_id = ?", UserID)
	return os
}

// 获取当前用户所有未支付订单
func GetOrdersByUserNotPay(UserID any) []entity.Order {
	var os []entity.Order
	Db.Find(&os, "user_id = ? and status = ?", UserID, false)
	return os
}

// 获取当前用户所有已支付订单
func GetOrdersByUserPaid(UserID any) []entity.Order {
	var os []entity.Order
	Db.Where("user_id = ? and status = ?", UserID, true).Scan(&os)
	return os
}

// 创建订单
func CreateOrder(o entity.Order) error {
	o.Time = time.Now().UnixMilli()
	return Db.Create(&o).Error
}

// 支付订单
func PayOrder(id any) error {
	return Db.Model(&entity.Order{}).Where("id = ?", id).Update("status", 0).Error
}

func DeleteOrder(id any) error {
	return Db.Delete(&entity.Order{}, id).Error
}

func AbortOrder(id any) error {
	fmt.Println(id)
	return Db.Model(&entity.Order{}).Where("id = ?", id).Update("status", 1).Error
}
