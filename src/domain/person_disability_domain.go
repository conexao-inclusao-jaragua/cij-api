package domain

import "cij_api/src/model"

type PersonDisabilityRepo interface {
	GetPersonDisabilities(personId int) ([]model.PersonDisability, error)
	GetDisabilityById(disabilityId int) (model.Disability, error)
	UpsertPersonDisability(personDisability model.PersonDisability) error
	ClearPersonDisability(personId int) error
}
