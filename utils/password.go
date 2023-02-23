package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

func ComparePassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

// 6位随机验证码
func RandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(rndCode)
	return rndCode
}
