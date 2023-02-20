package entity

import "gorm.io/gorm"

type ChargeStation struct {
	gorm.Model
	StationName string `gorm:"comment:充电站名称;type:varchar(100)" json:"stationName"`
	Area        string `gorm:"comment:所属校区;type:varchar(100)" json:"area"`
	// 定位
	AreaDescribe string  `gorm:"comment:地点描述;type:varchar(100)" json:"areaDescribe"`
	OpenTime     string  `gorm:"comment:开放时间段;type:varchar(100)" json:"openTime"`
	Price        float32 `gorm:"comment:单价" json:"price"`
	Coordinate   `gorm:"embedded"`
	// 详细信息
	ChargeStationDetails `gorm:"embedded"`
	// 充电桩
	ChargePoints []ChargePoint `gorm:"foreignKey:StationID"`
}

type ChargeStationDetails struct {
	Total int `gorm:"comment:充电桩总数" json:"total"`
	Spare int `gorm:"comment:空闲充电桩;" json:"spare"`
	Busy  int `gorm:"comment:忙碌充电桩;" json:"busy"`
}

type Coordinate struct {
	Latitude  float64 `gorm:"comment:纬度" json:"latitude"`
	Longitude float64 `gorm:"comment:经度" json:"longitude"`
}

// TableName 修改表名映射
func (ChargeStation) TableName() string {
	return "charge_stations"
}
