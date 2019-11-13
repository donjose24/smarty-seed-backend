package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/services/unionbank"
	"github.com/jmramos02/smarty-seed-backend/config"
	"time"
)

//TODO: lots of duplicate code. reuse some stuff.
type Claims struct {
	User models.User
	jwt.StandardClaims
}

type PledgeClaim struct {
	Pledge unionbank.GenerateUnionBankURLRequest
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

func DecodeUserInfo(token string) (models.User, error) {
	appKey := config.GetApplicationKey()
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	return claims.User, nil
}

func EncodePledge(pledge unionbank.GenerateUnionBankURLRequest) string {
	appKey := config.GetApplicationKey()
	// 1 hour expiration
	expirationTime := time.Now().Add(1 * time.Hour)
	claim := PledgeClaim{
		Pledge: pledge,
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

func DecodePledge(token string) (unionbank.GenerateUnionBankURLRequest, error) {
	appKey := config.GetApplicationKey()
	pledge := &PledgeClaim{}

	_, err := jwt.ParseWithClaims(token, pledge, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	if err != nil {
		return unionbank.GenerateUnionBankURLRequest{}, errors.New(err.Error())
	}

	return pledge.Pledge, nil
}
