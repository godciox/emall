package utils

import (
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logger.Debug(err)
	}
	return string(hash)
	//return string(pwd)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logger.Debug(err)
		return false
	}
	return true
}
