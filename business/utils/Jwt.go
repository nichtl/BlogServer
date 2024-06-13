package utils

import (
	"blogServe/business/global"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Account  string `json:"account"`
	jwt.StandardClaims
}

func GenerateToken(username, password, account string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		account,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
