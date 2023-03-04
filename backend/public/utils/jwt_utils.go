package utils

import (
	"fmt"
	"time"

	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string
	ID       string
	jwt.RegisteredClaims
}

func GenerateJWT(user *dao.User, secretKey string, expiredTime uint8) (string, error) {

	expirationTime := time.Now().Add(time.Duration(expiredTime) * time.Minute)
	claims := &Claims{
		Username: user.Username,
		ID:       fmt.Sprint(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string, tokenSecret string) (bool, *Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(tokenSecret), nil
	})

	if err != nil {
		return false, nil, err
	}

	if !token.Valid {
		return false, nil, nil
	}

	return true, claims, nil
}
