package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Time    int64  `gorm:"comment:订单发起时间;type:varchar(100)" json:"time"`
	Status  int    `gorm:"comment:支付状态，0已支付，1已取消，2未支付" json:"status"`
	Amount  string `gorm:"comment:订单金额;type:varchar(100)" json:"amount"`
	Charge  string `gorm:"comment:充电时长;type:varchar(100)" json:"charge"`
	Type    string `gorm:"comment:充电类型;type:varchar(100)" json:"type"`
	UserID  string `gorm:"comment:订单用户;type:varchar(100)" json:"userID"`
	PointID string `gorm:"comment:订单充电桩;type:varchar(100)" json:"pointID"`
}

// TableName 修改表名映射
func (Order) TableName() string {
	return "orders"
}
