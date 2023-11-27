package domain

import "cij_api/src/model"

type UserRepo interface {
	CreateUser(createUser model.User) (int, error)
	ListUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user model.User, userId int) error
	DeleteUser(id int) error
	DeleteUserPermanent(id int) error
}
