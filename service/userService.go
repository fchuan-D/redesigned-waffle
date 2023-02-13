package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"soft-pro/conf"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/middleware/redis"
	"soft-pro/resp"
	"soft-pro/utils"
	"strconv"
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
	// 判断用户名是否重复
	n := dao.CheckUserByName(u.UserName)
	if n != 0 {
		return errors.New(resp.NameExistErrorMsg)
	}

	return nil
}

// 校验登录数据
func CheckLoginUser(phone string, password string) (entity.User, error) {
	u, err := GetUserByPhone(phone)
	// 该手机号未注册
	if err != nil {
		return u, err
	}
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return u, errors.New(resp.LoginCheckErrorMsg)
	}

	return u, nil
}

func GetUserByID(id any) (entity.User, error) {
	u := dao.GetUserByID(id)
	if u.ID == 0 {
		return u, errors.New(resp.UserNotExistErrorMsg)
	}
	return u, nil
}

func GetUserByPhone(phone string) (entity.User, error) {
	u := dao.GetUserByPhone(phone)
	if u.ID == 0 {
		return u, errors.New(resp.PhoneNotExistErrorMsg)
	}
	return u, nil
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

// Redis缓存
func SaveBufferToRd(token string, u entity.User) {
	expireTime := time.Now().Add(time.Duration(conf.GetConfig().JwtAccessAge) * time.Minute).Unix()
	et := time.Unix(expireTime, 0)
	// key:token - value:User.ID
	redis.GetClient().Set(token, u.ID, et.Sub(time.Now()))
	// key:User.ID - value:User
	redis.GetClient().Set(strconv.Itoa(int(u.ID)), &u, 0)
}

// 获取手机号验证码
func SendPhoneCode(phone string) {
	// 获取随机验证码
	code := utils.RandCode()
	// redis记录
	redis.SetMini(phone, code, 5)
	// ToDo:异步发送验证码
	fmt.Println(phone, code)
}
