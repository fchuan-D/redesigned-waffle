package dao

import "soft-pro/entity"

func GetUserByID(id any) entity.User {
	var u entity.User
	Db.Find(&u, id)
	return u
}

func GetUserByPhone(phone string) entity.User {
	var u entity.User
	Db.Where("phone = ?", phone).First(&u)
	return u
}

func CheckUserByPhone(phone string) int64 {
	res := Db.Where("phone = ?", phone).First(&entity.User{})
	return res.RowsAffected
}

func InsertUser(u entity.User) {
	Db.Create(&u)
}
