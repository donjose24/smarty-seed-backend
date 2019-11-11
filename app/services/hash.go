package services

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		panic("Failed to hash string.")
	}

	return string(hash)
}

func CompareToHash(hash string, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err
}
