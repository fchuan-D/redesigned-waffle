package dao

import "soft-pro/entity"

func GetUserByID(id string) entity.User {
	var u entity.User
	Db.First(&u, id)
	return u
}
