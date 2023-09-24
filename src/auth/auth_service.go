package auth

import (
	"cij_api/src/config"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepo domain.UserRepo
}

func NewAuthService(userRepo domain.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func getSecretKey() ([]byte, error) {
	loadConfig, err := config.LoadConfig("../")
	if err != nil {
		return nil, errors.New("error on get secret key from .env")
	}

	return []byte(loadConfig.SecretKey), nil
}

func (s *AuthService) GenerateToken(user model.User) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", errors.New("error on get secret key from .env")
	}

	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"username":  user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Authenticate(credentials model.Credentials) (model.User, error) {
	user, err := s.userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		return user, errors.New("failed to get user by email")
	}

	if user.Email == "" {
		return user, errors.New("user with this email not found")
	}

	if !user.ValidatePassword(credentials.Password) {
		return user, errors.New("email/password incorrects")
	}

	return user, nil
}
