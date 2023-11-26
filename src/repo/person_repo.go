package repo

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) domain.PersonRepo {
	return &personRepo{
		db: db,
	}
}

func (n *personRepo) CreatePerson(createPerson model.Person) error {
	if err := n.db.Create(&createPerson).Error; err != nil {
		return errors.New("failed to create person")
	}

	return nil
}

func (n *personRepo) ListPeople() ([]model.Person, error) {
	var people []model.Person

	err := n.db.Model(model.Person{}).Preload("User").Find(&people).Error
	if err != nil {
		return people, errors.New("error on list people from database")
	}

	return people, nil
}

func (n *personRepo) GetPersonByUserId(userId int) (model.Person, error) {
	var person model.Person

	err := n.db.Model(model.Person{}).Preload("User").Where("user_id = ?", userId).Find(&person).Error
	if err != nil {
		return person, errors.New("failed to get the person")
	}

	return person, nil
}
