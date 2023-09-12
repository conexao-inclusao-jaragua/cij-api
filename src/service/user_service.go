package service

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(userRepo domain.UserRepo) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo domain.UserRepo
}

func (n *userService) CreateUser(createUser model.User) error {
	hashedPassword, err := encryptPassword(createUser.Password)
	if err != nil {
		return errors.New("error on encrypt user password")
	}

	createUser.Password = hashedPassword

	err = n.userRepo.CreateUser(createUser)

	if err != nil {
		return errors.New("error on create user")
	}

	return nil
}

func encryptPassword(password string) (string, error) {
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
