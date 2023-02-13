package test

import (
	"fmt"
	"soft-pro/conf"
	"soft-pro/dao"
	"testing"
)

func TestGetOrdersByUser(t *testing.T) {
	conf.InitConfig(".")
	dao.Init()

	o := dao.GetOrdersByUserID(1)
	fmt.Println(o)
	os := dao.GetOrdersByUserPaid(1)
	fmt.Println(os)
	os = dao.GetOrdersByUserNotPay(1)
	fmt.Println(o)
}
