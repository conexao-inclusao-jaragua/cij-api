package domain

import "cij_api/src/model"

type UserRepo interface {
	CreateUser(createUser model.User) error
	ListUsers() ([]model.User, error)
}

type UserService interface {
	CreateUser(createUser model.User) error
}
