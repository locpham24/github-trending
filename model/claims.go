package model

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtCustomClaims struct {
	UserId string
	Role   string
	jwt.StandardClaims
}

func GetToken(u User) (string, error) {
	claims := &JwtCustomClaims{
		UserId: u.UserId,
		Role:   u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(os.Getenv("JwtSecretKey")))
	if err != nil {
		return "", err
	}

	return result, nil
}
