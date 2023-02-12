package entity

import "gorm.io/gorm"

type ChargePoint struct {
	gorm.Model
	InUse   bool `gorm:"comment:是否被占用"`
	Remains uint `gorm:"comment:剩余充电时长,单位分钟"`
}

// TableName 修改表名映射
func (ChargePoint) TableName() string {
	return "charge_points"
}
