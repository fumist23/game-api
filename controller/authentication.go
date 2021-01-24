package controller

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenWithName(name string) (string, error) {

	// 鍵となる文字列(ここではセキュリティを求めないので全部この鍵を用いる)
	secretKey := ""

	// トークンはjwtを使用
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": name,
	})
	fmt.Printf("token is: %v", token)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("failed to sign token: &v", err)
		return "", err
	}

	fmt.Printf("tokenString: %v", tokenString)

	return tokenString, nil

}
