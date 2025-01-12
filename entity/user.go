package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	UserName    string  `gorm:"comment:用户名;type:varchar(20);not null;unique" json:"userName"`
	Phone       string  `gorm:"comment:手机号;type:varchar(20);not null;unique" json:"telephone"`
	Password    string  `gorm:"comment:密码;size:255;not null" json:"-"`
	Role        string  `gorm:"comment:用户角色;type:varchar(20)" json:"role"`
	Bal         float64 `gorm:"comment:余额;column:balance;default:0" json:"balance"`
	Ban         bool    `gorm:"comment:是否已开通余额;default:0" json:"ban"`
	TotalCharge float64 `gorm:"comment:总充电量;column:total_charge;default:0" json:"totalCharge"`
	TotalPay    float64 `gorm:"comment:总消费;column:total_pay;default:0" json:"totalPay"`
	TotalBal    float64 `gorm:"comment:总充值;column:total_bal;default:0" json:"totalBal"`
	Orders      []Order `gorm:"foreignKey:UserID" json:"orders"`
}

// 序列化
func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// 反序列化
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)

}

// 获取当前余额
func (u User) getBal() float64 {
	return u.Bal
}

// 更新当前余额
func (u *User) updateBal(inc float64) {
	u.Bal += inc
}

// TableName 修改表名映射
func (User) TableName() string {
	return "users"
}
