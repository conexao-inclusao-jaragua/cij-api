package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error on encrypt password")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		return "", errors.New("error on encrypt password")
	}

	return string(hashedPassword), nil
}
