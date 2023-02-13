package dao

import "soft-pro/entity"

func GetChargeStationByID(id any) entity.ChargeStation {
	var cs entity.ChargeStation
	Db.Find(&cs, "id = ?", id)
	return cs
}

func GetChargeStationByArea(area string) entity.ChargeStation {
	var cs entity.ChargeStation
	Db.Find(&cs, "area = ?", area)
	return cs
}

func GetTotal(id any) int {
	l := len(GetTotalChargePoint(id))
	return l
}

func GetSpare(id any) int {
	l := len(GetSpareChargePoint(id))
	return l
}

func GetBusy(id any) int {
	l := len(GetBusyChargePoint(id))
	return l
}
