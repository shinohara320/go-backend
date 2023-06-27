package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.At(expirationTime),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseJwt(cookie string) (uint, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid or expired token")
	}
	claims := token.Claims.(*jwt.StandardClaims)
	Id, err := strconv.ParseUint(claims.Issuer, 10, 32)
	if err != nil {
		return 0, errors.New("invalid user id in token")
	}
	return uint(Id), nil
}
