package repo

import (
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type PersonRepo interface {
	CreatePerson(createPerson model.Person) error
	ListPeople() ([]model.Person, error)
	GetPersonById(personId int) (model.Person, error)
	GetPersonByUserId(userId int) (model.Person, error)
	GetPersonByCpf(cpf string) (model.Person, error)
	GetPersonDisabilities(personId int) ([]model.Disability, error)
	GetDisabilityById(disabilityId int) (model.Disability, error)
	UpdatePerson(person model.Person, personId int) error
	UpsertPersonDisability(disability model.Disability, personId int) error
	ClearPersonDisability(personId int) error
	DeletePerson(personId int) error
}

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) PersonRepo {
	return &personRepo{
		db: db,
	}
}

func (n *personRepo) CreatePerson(createPerson model.Person) (int, error) {
	if err := n.db.Create(&createPerson).Error; err != nil {
		return 0, errors.New("failed to create person")
	}

	return createPerson.Id, nil
}

func (n *personRepo) ListPeople() ([]model.Person, error) {
	var people []model.Person

	err := n.db.Model(model.Person{}).Preload("User").Find(&people).Error
	if err != nil {
		return people, errors.New("error on list people from database")
	}

	return people, nil
}

func (n *personRepo) GetPersonById(personId int) (model.Person, error) {
	var person model.Person

	err := n.db.Model(model.Person{}).Preload("User").Preload("Address").Where("id = ?", personId).Find(&person).Error
	if err != nil {
		return person, errors.New("failed to get the person")
	}

	return person, nil
}

func (n *personRepo) GetPersonByUserId(userId int) (model.Person, error) {
	var person model.Person

	err := n.db.Model(model.Person{}).Preload("User").Where("user_id = ?", userId).Find(&person).Error
	if err != nil {
		return person, errors.New("failed to get the person")
	}

	return person, nil
}

func (n *personRepo) GetPersonByCpf(cpf string) (model.Person, error) {
	var person model.Person

	err := n.db.Model(model.Person{}).Preload("User").Where("cpf = ?", cpf).Find(&person).Error
	if err != nil {
		return person, errors.New("failed to get the person")
	}

	return person, nil
}

func (n *personRepo) UpdatePerson(person model.Person, personId int) error {
	if err := n.db.Model(model.Person{}).Where("id = ?", personId).Updates(person).Error; err != nil {
		return errors.New("failed to update the person")
	}

	return nil
}

func (n *personRepo) DeletePerson(personId int) error {
	if err := n.db.Model(model.Person{}).Where("id = ?", personId).Unscoped().Delete(&model.Person{}).Error; err != nil {
		return errors.New("failed to delete the person")
	}

	return nil
}
