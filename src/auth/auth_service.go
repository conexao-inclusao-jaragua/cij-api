package auth

import (
	"cij_api/src/config"
	"cij_api/src/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func getSecretKey() ([]byte, error) {
	loadConfig, err := config.LoadConfig("../")
	if err != nil {
		return nil, errors.New("error on get secret key from .env")
	}

	return []byte(loadConfig.SecretKey), nil
}

func GenerateToken(user model.User) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", errors.New("error on get secret key from .env")
	}

	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = user.Name

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
