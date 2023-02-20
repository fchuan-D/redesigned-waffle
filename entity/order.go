package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Status  bool   `gorm:"comment:支付状态"`
	Amount  string `gorm:"comment:订单金额"`
	Charge  string `gorm:"comment:充电时长"`
	Type    string `gorm:"comment:充电类型"`
	UserID  string `gorm:"comment:订单用户" json:"userID"`
	PointID string `gorm:"comment:订单充电桩" json:"pointID"`
}

// TableName 修改表名映射
func (Order) TableName() string {
	return "orders"
}
