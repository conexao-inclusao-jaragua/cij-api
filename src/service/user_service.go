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

func (s *userService) ListUsers() ([]model.User, error) {
	users, err := s.userRepo.ListUsers()
	if err != nil {
		return users, errors.New("failed to list users")
	}

	return users, nil
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

func (n *userService) GetUserByEmail(email string) (model.User, error) {
	user, err := n.userRepo.GetUserByEmail(email)
	if err != nil {
		return user, errors.New("failed to get user by email")
	}

	return user, nil
}
