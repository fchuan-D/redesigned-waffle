package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"soft-pro/conf"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/resp"
	"time"
)

// 校验注册数据
func CheckRegisterUser(u entity.User) error {
	// 格式校验
	if len(u.UserName) == 0 || len(u.UserName) > 20 {
		return errors.New(resp.NameCheckErrorMsg)
	}
	if len(u.Phone) != 11 {
		return errors.New(resp.PhoneCheckErrorMsg)
	}
	if len(u.Password) < 6 || len(u.Password) > 20 {
		return errors.New(resp.PwdCheckErrorMsg)
	}

	// 判断手机号是否存在
	i := dao.CheckUserByPhone(u.Phone)
	if i != 0 {
		return errors.New(resp.PhoneExistErrorMsg)
	}

	return nil
}

// 校验登录数据
func CheckLoginUser(phone string, password string) (entity.User, error) {
	u := GetUserByPhone(phone)
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return u, errors.New(resp.LoginCheckErrorMsg)
	}

	return u, nil
}

func GetUserByID(id any) entity.User {
	return dao.GetUserByID(id)
}

func GetUserByPhone(phone string) entity.User {
	return dao.GetUserByPhone(phone)
}

func InsertUser(u entity.User) error {
	// 密码加密存储
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密错误")
	}
	u.Password = string(hasedPassword)
	dao.InsertUser(u)
	return nil
}

// Redis缓存 	key:token - value:User
func SaveBufferToRd(token string, u entity.User) {
	expireTime := time.Now().Add(time.Duration(conf.GetConfig().JwtAccessAge) * time.Minute).Unix()
	et := time.Unix(expireTime, 0)
	dao.Rd.Set(token, &u, et.Sub(time.Now()))
}
