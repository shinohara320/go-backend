package util

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer":    issuer,
		"ExpiresAt": expirationTime.Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
