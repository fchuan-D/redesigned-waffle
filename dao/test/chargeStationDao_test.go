package test

import (
	"fmt"
	"soft-pro/conf"
	"soft-pro/dao"
	"testing"
)

func TestChargeStation_TableName(t *testing.T) {
	conf.InitConfig(".")
	dao.Init()

	fmt.Println(dao.GetSpare(1))
}
