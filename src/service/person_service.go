package service

import (
	"cij_api/src/auth"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

type personService struct {
	personRepo domain.PersonRepo
	userRepo   domain.UserRepo
}

func NewPersonService(personRepo domain.PersonRepo, userRepo domain.UserRepo) domain.PersonService {
	return &personService{
		personRepo: personRepo,
		userRepo:   userRepo,
	}
}

func (s *personService) ListPeople() ([]model.PersonResponse, error) {
	peopleResponse := []model.PersonResponse{}

	people, err := s.personRepo.ListPeople()
	if err != nil {
		return peopleResponse, errors.New("failed to list people")
	}

	for _, person := range people {
		user, err := s.userRepo.GetUserById(person.User.Id)
		if err != nil {
			return peopleResponse, errors.New("failed to get user")
		}

		peopleResponse = append(peopleResponse, person.ToResponse(user))
	}

	return peopleResponse, nil
}

func (n *personService) CreatePerson(createPerson model.PersonRequest) error {
	userInfo := createPerson.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return errors.New("error on encrypt user password")
	}

	userInfo.Password = hashedPassword
	userInfo.RoleId = 1 // 1 is the id of the role "person"

	userId, err := n.userRepo.CreateUser(userInfo)
	if err != nil {
		return errors.New("error on create user")
	}

	createPerson.User.Id = userId

	err = n.personRepo.CreatePerson(createPerson.ToPerson())
	if err != nil {
		n.userRepo.DeleteUser(userId)

		return errors.New("error on create person")
	}

	return nil
}
