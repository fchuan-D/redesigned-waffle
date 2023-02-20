package entity

import "gorm.io/gorm"

type ChargePoint struct {
	gorm.Model
	InUse     bool    `gorm:"comment:是否被占用;default:0" json:"inUse"`
	Remains   uint    `gorm:"comment:剩余充电时长,单位分钟;default:0" json:"remains"`
	StationID uint    `gorm:"comment:所在充电站ID;not null" json:"stationID"`
	Orders    []Order `gorm:"foreignKey:PointID" json:"orders"`
}

// TableName 修改表名映射
func (ChargePoint) TableName() string {
	return "charge_points"
}
