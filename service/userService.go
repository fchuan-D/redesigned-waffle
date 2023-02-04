package service

import (
	"soft-pro/dao"
	"soft-pro/entity"
)

func Search(id string) entity.User {
	return dao.GetUserByID(id)
}
