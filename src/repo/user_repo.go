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
