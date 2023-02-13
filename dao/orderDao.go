package dao

import "soft-pro/entity"

// 根据 OrderID获取订单数据
func GetOrderByID(id any) entity.Oreder {
	var o entity.Oreder
	Db.Find(&o, id)
	return o
}

// 获取当前用户所有订单
func GetOrdersByUserID(UserID any) []entity.Oreder {
	var os []entity.Oreder
	Db.Find(&os, "user_id = ?", UserID)
	return os
}

// 获取当前用户所有未支付订单
func GetOrdersByUserNotPay(UserID any) []entity.Oreder {
	var os []entity.Oreder
	Db.Find(&os, "user_id = ? and status = ?", UserID, false)
	return os
}

// 获取当前用户所有已支付订单
func GetOrdersByUserPaid(UserID any) []entity.Oreder {
	var os []entity.Oreder
	Db.Where("user_id = ? and status = ?", UserID, true).Scan(&os)
	return os
}
