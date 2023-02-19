package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status bool `gorm:"comment:支付状态"`
	Amount int  `gorm:"comment:订单金额"`
	Charge uint `gorm:"comment:充电时长"`
	UserID uint `gorm:"comment:订单用户"`
}

// TableName 修改表名映射
func (Order) TableName() string {
	return "orders"
}
