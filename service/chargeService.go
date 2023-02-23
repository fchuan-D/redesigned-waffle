package service

import (
	"errors"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/resp"
)

// 获取充电站列表
func GetStationList() ([]entity.ChargeStation, error) {
	css := dao.GetAllChargeStation()
	if css[0].ID == 0 {
		return css, errors.New(resp.NotFindMsg)
	}
	return css, nil
}

// 获取充电站下所有充电桩
func GetPointList(StationID string) ([]entity.ChargePoint, error) {
	cps := dao.GetTotalChargePoint(StationID)
	if cps[0].ID == 0 {
		return cps, errors.New(resp.NotFindMsg)
	}
	return cps, nil
}
