package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/config"
	"time"
)

type Claims struct {
	User models.User
	jwt.StandardClaims
}

func EncodeUserInfo(user models.User) string {
	appKey := config.GetApplicationKey()
	// 1 hour expiration
	expirationTime := time.Now().Add(1 * time.Hour)
	claim := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(appKey))

	if err != nil {
		panic("Signing Failed. Please check application key")
	}

	return tokenString
}
