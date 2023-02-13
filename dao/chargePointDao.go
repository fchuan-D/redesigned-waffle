package dao

import "soft-pro/entity"

func GetChargePointByID(id any) entity.ChargePoint {
	var cp entity.ChargePoint
	Db.Find(&cp, id)
	return cp
}

// 获取当前充电站下所有空闲充电桩
func GetSpareChargePoint(StationID any) []entity.ChargePoint {
	var cps []entity.ChargePoint
	Db.Find(&cps, "station_id = ? and in_use = ?", StationID, false)
	return cps
}

// 获取当前充电站下所有忙碌充电桩
func GetBusyChargePoint(StationID any) []entity.ChargePoint {
	var cps []entity.ChargePoint
	Db.Find(&cps, "station_id = ? and in_use = ?", StationID, true)
	return cps
}

// 获取当前充电站下所有的充电桩
func GetTotalChargePoint(StationID any) []entity.ChargePoint {
	var cps []entity.ChargePoint
	Db.Find(&cps, "station_id = ?", StationID)
	return cps
}

func InsertChargePoint(cp entity.ChargePoint) {
	Db.Create(&cp)
}
