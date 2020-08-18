package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string, authSecretKey string) (*jwt.Token, error) {
	if token == "" {
		return nil, errors.New("Token not found")
	}

	strArr := strings.Split(token, " ")
	if len(strArr) != 2 {
		return nil, errors.New("Token not found")
	}

	encodedStr := strArr[1]
	tokenData, err := jwt.Parse(encodedStr, func(data *jwt.Token) (interface{}, error) {
		if _, ok := data.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", data.Header["alg"])
		}

		return []byte(authSecretKey), nil
	})
	return tokenData, err
}
