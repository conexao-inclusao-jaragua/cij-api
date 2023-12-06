package domain

import (
	"cij_api/src/model"
	"cij_api/src/utils"
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

type PersonService interface {
	CreatePerson(createPerson model.PersonRequest) utils.Error
	ListPeople() ([]model.PersonResponse, utils.Error)
	GetPersonByUserId(userId int) (model.Person, utils.Error)
	GetPersonById(personId int) (model.Person, utils.Error)
	GetPersonByCpf(cpf string) (model.Person, utils.Error)
	GetUserByEmail(email string) (model.User, utils.Error)
	GetDisabilityById(disabilityId int) (model.Disability, utils.Error)
	UpdatePerson(person model.PersonRequest, personId int) utils.Error
	UpdatePersonAddress(address model.AddressRequest, personId int) utils.Error
	UpdatePersonDisabilities(disabilities []int, personId int) utils.Error
	DeletePerson(personId int) utils.Error
}
