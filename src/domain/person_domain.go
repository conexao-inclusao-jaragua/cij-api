package domain

import "cij_api/src/model"

type PersonRepo interface {
	CreatePerson(createPerson model.Person) error
	ListPeople() ([]model.Person, error)
	GetPersonByUserId(userId int) (model.Person, error)
}

type PersonService interface {
	CreatePerson(createPerson model.PersonRequest) error
	ListPeople() ([]model.PersonResponse, error)
	GetPersonByUserId(userId int) (model.Person, error)
}
