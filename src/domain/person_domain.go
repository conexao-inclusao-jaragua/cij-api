package domain

import "cij_api/src/model"

type PersonRepo interface {
	CreatePerson(createPerson model.Person) error
	ListPeople() ([]model.Person, error)
	GetPersonById(personId int) (model.Person, error)
	GetPersonByUserId(userId int) (model.Person, error)
	GetPersonDisabilities(personId int) ([]model.Disability, error)
	UpdatePerson(person model.Person, personId int) error
	UpsertPersonDisability(disability model.Disability, personId int) error
	ClearPersonDisability(personId int) error
	DeletePerson(personId int) error
}

type PersonService interface {
	CreatePerson(createPerson model.PersonRequest) error
	ListPeople() ([]model.PersonResponse, error)
	GetPersonByUserId(userId int) (model.Person, error)
	UpdatePerson(person model.PersonRequest, personId int) error
	UpdatePersonAddress(address model.AddressRequest, personId int) error
	UpdatePersonDisabilities(disabilities []int, personId int) error
	DeletePerson(personId int) error
}
