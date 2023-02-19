package entity

import "gorm.io/gorm"

type ChargeStation struct {
	gorm.Model
	Area string `gorm:"comment:所属校区"`
	// 定位
	Coordinate `gorm:"embedded"`
	// 详细信息
	ChargeStationDetails `gorm:"embedded"`
	// 充电桩
	ChargePoints []ChargePoint `gorm:"foreignKey:StationID"`
}

type ChargeStationDetails struct {
	Total int `gorm:"comment:充电桩总数"`
	Spare int `gorm:"comment:空闲充电桩;"`
	Busy  int `gorm:"comment:忙碌充电桩;"`
}

type Coordinate struct {
	Latitude  float64 `gorm:"comment:纬度"`
	Longitude float64 `gorm:"comment:经度"`
}

// TableName 修改表名映射
func (ChargeStation) TableName() string {
	return "charge_stations"
}
