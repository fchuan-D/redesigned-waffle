package test

import (
	"fmt"
	"soft-pro/conf"
	"soft-pro/dao"
	"testing"
)

func TestChargePoint_TableName(t *testing.T) {
	conf.InitConfig(".")
	dao.Init()

	cps := dao.GetTotalChargePoint(1)
	for i, cp := range cps {
		fmt.Println(i, cp)
	}

}
