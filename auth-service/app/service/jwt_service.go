package service

import (
	"auth-service/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int) (string, error) {

	expiredMinute, _ := strconv.Atoi(
		config.GetEnv("ACCESS_TOKEN_EXPIRED"),
	)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(
			time.Minute * time.Duration(expiredMinute),
		).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := []byte(config.GetEnv("JWT_SECRET"))

	return token.SignedString(secret)
}

func GenerateRefreshToken(userID int) (string, error) {

	expiredHour, _ := strconv.Atoi(
		config.GetEnv("REFRESH_TOKEN_EXPIRED"),
	)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(
			time.Hour * time.Duration(expiredHour),
		).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := []byte(config.GetEnv("JWT_SECRET"))

	return token.SignedString(secret)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	secret := []byte(config.GetEnv("JWT_SECRET"))

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}
