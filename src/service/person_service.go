package service

import (
	"cij_api/src/auth"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

type personService struct {
	personRepo  domain.PersonRepo
	userRepo    domain.UserRepo
	addressRepo domain.AddressRepo
}

func NewPersonService(personRepo domain.PersonRepo, userRepo domain.UserRepo, addressRepo domain.AddressRepo) domain.PersonService {
	return &personService{
		personRepo:  personRepo,
		userRepo:    userRepo,
		addressRepo: addressRepo,
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

		personResponse := person.ToResponse(user)

		if person.AddressId != nil {
			address, err := s.addressRepo.GetAddressById(*person.AddressId)
			if err != nil {
				return peopleResponse, errors.New("failed to get address")
			}

			if address.Id != 0 {
				var addressResponse model.AddressResponse
				addressResponse = address.ToResponse()
				personResponse.Address = &addressResponse
			}
		}

		disabilities, err := s.personRepo.GetPersonDisabilities(person.Id)
		if err != nil {
			return peopleResponse, errors.New("failed to get person disabilities")
		}

		if len(disabilities) > 0 {
			var disabilitiesResponse []model.DisabilityResponse

			for _, disability := range disabilities {
				disabilityResponse := disability.ToResponse()
				disabilitiesResponse = append(disabilitiesResponse, disabilityResponse)
			}

			personResponse.Disabilities = &disabilitiesResponse
		}

		peopleResponse = append(peopleResponse, personResponse)
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

	personInfo := createPerson.ToModel(userInfo)
	personInfo.UserId = userId

	err = n.personRepo.CreatePerson(personInfo)
	if err != nil {
		err = n.userRepo.DeleteUserPermanent(userId)
		if err != nil {
			return errors.New("error on delete user")
		}

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

	personInfo := updatePerson.ToModel(userInfo)

	err = n.personRepo.UpdatePerson(personInfo, personId)
	if err != nil {
		return errors.New("failed to update person")
	}

	return nil
}

func (n *personService) UpdatePersonAddress(updateAddress model.AddressRequest, personId int) error {
	addressInfo := updateAddress.ToModel()

	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return errors.New("failed to get person")
	}

	if *person.AddressId != 0 {
		addressInfo.Id = *person.AddressId
	}

	addressId, err := n.addressRepo.UpsertAddress(addressInfo)
	if err != nil {
		return errors.New("failed to upsert address")
	}

	person.AddressId = &addressId

	err = n.personRepo.UpdatePerson(person, personId)
	if err != nil {
		return errors.New("failed to update person")
	}

	return nil
}

func (n *personService) UpdatePersonDisabilities(disabilities []int, personId int) error {
	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return errors.New("failed to get person")
	}

	err = n.personRepo.ClearPersonDisability(personId)
	if err != nil {
		return errors.New("failed to clear person disability")
	}

	for _, disabilityId := range disabilities {
		disability := model.Disability{
			Id: disabilityId,
		}

		err = n.personRepo.UpsertPersonDisability(disability, personId)
		if err != nil {
			return errors.New("failed to upsert person disability")
		}
	}

	err = n.personRepo.UpdatePerson(person, personId)
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

	err = n.addressRepo.DeleteAddress(*person.AddressId)
	if err != nil {
		return errors.New("failed to delete address")
	}

	return nil
}
