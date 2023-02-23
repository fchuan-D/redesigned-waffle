package dao

import "soft-pro/entity"

func GetUserByID(id any) entity.User {
	var u entity.User
	Db.Preload("Orders").Find(&u, id)
	return u
}

func GetUserByPhone(phone string) entity.User {
	var u entity.User
	Db.Where("phone = ?", phone).First(&u)
	return u
}

func CheckUserByPhone(phone string, uid any) int64 {
	res := Db.Where("phone = ? and id != ?", phone, uid).First(&entity.User{})
	return res.RowsAffected
}

func CheckUserByName(name string, uid any) int64 {
	res := Db.Where("user_name = ? and id != ?", name, uid).First(&entity.User{})
	return res.RowsAffected
}

func InsertUser(u entity.User) error {
	return Db.Save(&u).Error
}

func UpdateBal(uid uint, bal float64) error {
	return Db.Model(&entity.User{}).Where("id = ?", uid).Update("balance", bal).Error
}
