package token

import (
	"fmt"
	_const "frontend/const"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(Phone string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	tokenexp := _const.JWTMaxAge
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix()
	//claims["exp"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()
	claims["phone"] = Phone
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(_const.TokenSecret))
	return tokenString
}

func CheckToken(tokenString string) (string, bool) {
	Phone := ""
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(_const.TokenSecret), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	Phone = claims["phone"].(string)
	//iat := int64(claims["iat"].(float64))
	exp := int64(claims["exp"].(float64))
	//begin := time.Unix(iat, 0)
	expT := time.Unix(exp, 0)
	if time.Now().Before(expT) {
		return "", false
	}
	return Phone, true
}
