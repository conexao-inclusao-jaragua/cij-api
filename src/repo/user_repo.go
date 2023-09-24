package repo

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (n *userRepo) CreateUser(createUser model.User) error {
	if err := n.db.Create(&createUser).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (n *userRepo) ListUsers() ([]model.User, error) {
	var users []model.User

	err := n.db.Model(model.User{}).Find(&users).Error
	if err != nil {
		return users, errors.New("error on list users from database")
	}

	return users, nil
}

func (n *userRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	err := n.db.Model(model.User{}).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, errors.New("failed to get the user")
	}

	return user, nil
}
