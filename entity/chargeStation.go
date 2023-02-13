package entity

import "gorm.io/gorm"

type ChargeStation struct {
	gorm.Model
	Coordinate `gorm:"embedded"`
	Area       string `gorm:"comment:所属校区"`
	Total      int    `gorm:"comment:充电桩总数"`
	Spare      int    `gorm:"comment:空闲充电桩;"`
	Busy       int    `gorm:"comment:忙碌充电桩;"`
}

// TableName 修改表名映射
func (ChargeStation) TableName() string {
	return "charge_stations"
}
