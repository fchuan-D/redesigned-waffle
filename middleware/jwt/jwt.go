package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"soft-pro/conf"
	"soft-pro/entity"
	"soft-pro/middleware/redis"
	"soft-pro/resp"
	"soft-pro/service"
	"strconv"
	"time"
)

type MyClaim struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateToken(id uint, name string) (string, error) {
	expireTime := time.Now().Add(time.Duration(conf.GetConfig().JwtAccessAge) * time.Minute)
	// 实例化claim
	claim := MyClaim{
		UserID: id,
		// 定义jwt自带的属性
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 生命周期
			Issuer:    name,              // 签发者
		},
	}

	// 通过 claim创建原始 token
	// 参数一：加密编码；参数二：使用的 claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 使用指定的密钥对原始 token进行签名，获得加密后的 token
	return token.SignedString([]byte(conf.GetConfig().JwtKey))
}

func ParseToken(tokenString string) (*MyClaim, error) {
	claim := &MyClaim{}

	// 解析 tokenString的固定格式，第三个参数 golang会自动生成
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetConfig().JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid { // 校验 token
		return claims, nil
	}
	return nil, errors.New(resp.TokenInValidErrorMsg)
}

func CheckBufferToken(token string, UserID uint) (entity.User, error) {
	rd := redis.GetClient()
	// 合法的 token,进一步校验是否过期
	// 根据 token拿到缓存中的 UserID
	var u entity.User
	var i uint
	err := rd.Get(token).Scan(&i)
	//校验 UserID是否一致
	if i != UserID || err != nil {
		// 缓存数据出错
		return u, errors.New(resp.TokenInValidErrorMsg)
	}
	// 根据 UserID获取缓存的 User信息
	// 缓存未命中再从mysql中查询并放入缓存
	if err := rd.Get(strconv.Itoa(int(i))).Scan(&u); err != nil {
		u, err := service.GetUserByID(i)
		if err != nil {
			return u, errors.New(err.Error())
		}
		// key:User.ID - value:User
		rd.Set(strconv.Itoa(int(u.ID)), &u, 0)
	}
	return u, nil
}
