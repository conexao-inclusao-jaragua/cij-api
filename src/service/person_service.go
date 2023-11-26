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

	personInfo := createPerson.ToPerson(userInfo)
	personInfo.UserId = userId

	err = n.personRepo.CreatePerson(personInfo)
	if err != nil {
		n.userRepo.DeleteUser(userId)

		return errors.New("error on create person")
	}

	return nil
}

func (n *personService) GetPersonByUserId(userId int) (model.Person, error) {
	person, err := n.personRepo.GetPersonByUserId(userId)
	if err != nil {
		return person, errors.New("failed to get person")
	}

	return person, nil
}

func (n *personService) UpdatePerson(updatePerson model.PersonRequest, personId int) error {
	userInfo := updatePerson.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return errors.New("error on encrypt user password")
	}

	userInfo.Password = hashedPassword

	err = n.userRepo.UpdateUser(userInfo, personId)
	if err != nil {
		return errors.New("failed to update user")
	}

	personInfo := updatePerson.ToPerson(userInfo)

	err = n.personRepo.UpdatePerson(personInfo, personId)
	if err != nil {
		return errors.New("failed to update person")
	}

	return nil
}

func (n *personService) DeletePerson(personId int) error {
	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return errors.New("failed to get person")
	}

	err = n.personRepo.DeletePerson(personId)
	if err != nil {
		return errors.New("failed to delete person")
	}

	err = n.userRepo.DeleteUser(person.UserId)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
