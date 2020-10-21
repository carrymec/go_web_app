package tool

import (
	"golang.org/x/crypto/bcrypt"
)

func AesEncrypt(password string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", true
	}
	return string(hash), false
}

//数据库密码+登录输入密码
func AesDecrypt(origin string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte(password)) //验证（对比）
	if err != nil {
		return false
	}
	return true
}
