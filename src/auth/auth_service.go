package auth

import (
	"cij_api/src/config"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo    domain.UserRepo
	companyRepo domain.CompanyRepo
}

func NewAuthService(userRepo domain.UserRepo, companyRepo domain.CompanyRepo) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		companyRepo: companyRepo,
	}
}

func getSecretKey() ([]byte, error) {
	loadConfig, err := config.LoadConfig("../")
	if err != nil {
		return nil, errors.New("error on get secret key from .env")
	}

	return []byte(loadConfig.SecretKey), nil
}

func (s *AuthService) GenerateToken(role string) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", errors.New("error on get secret key from .env")
	}

	claims := &jwt.MapClaims{
		"exp":  jwt.TimeFunc().Add(time.Minute * 10).Unix(),
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret, err := getSecretKey()
	if err != nil {
		return nil, errors.New("failed to get token")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

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

func (s *AuthService) Authenticate(credentials model.Credentials, role string) (model.User, model.Company, error) {
	var user model.User
	var company model.Company

	if role == "user" {
		user, err := s.userRepo.GetUserByEmail(credentials.Email)
		if err != nil {
			return user, company, errors.New("failed to get user by email")
		}

		if user.Email == "" {
			return user, company, errors.New("user with this email not found")
		}

		if !user.ValidatePassword(credentials.Password) {
			return user, company, errors.New("email/password incorrects")
		}

		return user, company, nil
	}

	if role == "company" {
		company, err := s.companyRepo.GetCompanyByEmail(credentials.Email)
		if err != nil {
			return user, company, errors.New("failed to get company by email")
		}

		if company.Email == "" {
			return user, company, errors.New("company with this email not found")
		}

		if !company.ValidatePassword(credentials.Password) {
			return user, company, errors.New("email/password incorrects")
		}

		return user, company, nil
	}

	return user, company, errors.New("role not found")
}
