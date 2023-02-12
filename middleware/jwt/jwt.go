package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"soft-pro/conf"
	"soft-pro/resp"
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
