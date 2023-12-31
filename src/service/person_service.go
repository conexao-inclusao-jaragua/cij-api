package service

import (
	"cij_api/src/auth"
	"cij_api/src/domain"
	"cij_api/src/model"
	"cij_api/src/utils"
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

func (s *personService) ListPeople() ([]model.PersonResponse, utils.Error) {
	peopleResponse := []model.PersonResponse{}

	people, err := s.personRepo.ListPeople()
	if err != nil {
		return peopleResponse, utils.FailedToListPeople
	}

	for _, person := range people {
		user, err := s.userRepo.GetUserById(person.User.Id)
		if err != nil {
			return peopleResponse, utils.FailedToGetUser
		}

		personResponse := person.ToResponse(user)

		if person.AddressId != nil {
			address, err := s.addressRepo.GetAddressById(*person.AddressId)
			if err != nil {
				return peopleResponse, utils.FailedToGetAddress
			}

			if address.Id != 0 {
				var addressResponse model.AddressResponse
				addressResponse = address.ToResponse()
				personResponse.Address = &addressResponse
			}
		}

		disabilities, err := s.personRepo.GetPersonDisabilities(person.Id)
		if err != nil {
			return peopleResponse, utils.FailedToGetDisability
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

	return peopleResponse, utils.Error{}
}

func (n *personService) CreatePerson(createPerson model.PersonRequest) utils.Error {
	userInfo := createPerson.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return utils.FailedToEncryptPassword
	}

	userInfo.Password = hashedPassword
	userInfo.RoleId = 1 // 1 is the id of the role "person"

	userId, err := n.userRepo.CreateUser(userInfo)
	if err != nil {
		return utils.FailedToCreateUser
	}

	personInfo := createPerson.ToModel(userInfo)
	personInfo.UserId = userId

	err = n.personRepo.CreatePerson(personInfo)
	if err != nil {
		err = n.userRepo.DeleteUserPermanent(userId)
		if err != nil {
			return utils.FailedToDeleteUser
		}

		return utils.FailedToCreatePerson
	}

	return utils.Error{}
}

func (n *personService) GetPersonByUserId(userId int) (model.Person, utils.Error) {
	person, err := n.personRepo.GetPersonByUserId(userId)
	if err != nil {
		return person, utils.FailedToGetPerson
	}

	return person, utils.Error{}
}

func (n *personService) GetPersonById(personId int) (model.Person, utils.Error) {
	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return person, utils.FailedToGetPerson
	}

	return person, utils.Error{}
}

func (n *personService) GetPersonByCpf(cpf string) (model.Person, utils.Error) {
	person, err := n.personRepo.GetPersonByCpf(cpf)
	if err != nil {
		return person, utils.FailedToGetPerson
	}

	return person, utils.Error{}
}

func (n *personService) GetUserByEmail(email string) (model.User, utils.Error) {
	user, err := n.userRepo.GetUserByEmail(email)
	if err != nil {
		return user, utils.FailedToGetUser
	}

	return user, utils.Error{}
}

func (n *personService) UpdatePerson(updatePerson model.PersonRequest, personId int) utils.Error {
	userInfo := updatePerson.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return utils.FailedToEncryptPassword
	}

	userInfo.Password = hashedPassword

	err = n.userRepo.UpdateUser(userInfo, personId)
	if err != nil {
		return utils.FailedToUpdateUser
	}

	personInfo := updatePerson.ToModel(userInfo)

	err = n.personRepo.UpdatePerson(personInfo, personId)
	if err != nil {
		return utils.FailedToUpdatePerson
	}

	return utils.Error{}
}

func (n *personService) UpdatePersonAddress(updateAddress model.AddressRequest, personId int) utils.Error {
	addressInfo := updateAddress.ToModel()

	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return utils.FailedToGetPerson
	}

	if person.AddressId != nil {
		addressInfo.Id = *person.AddressId
	}

	addressId, err := n.addressRepo.UpsertAddress(addressInfo)
	if err != nil {
		return utils.FailedToUpsertAddress
	}

	person.AddressId = &addressId

	err = n.personRepo.UpdatePerson(person, personId)
	if err != nil {
		return utils.FailedToUpdatePerson
	}

	return utils.Error{}
}

func (n *personService) GetDisabilityById(id int) (model.Disability, utils.Error) {
	disability, err := n.personRepo.GetDisabilityById(id)
	if err != nil {
		return disability, utils.FailedToGetDisability
	}

	return disability, utils.Error{}
}

func (n *personService) UpdatePersonDisabilities(disabilities []int, personId int) utils.Error {
	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return utils.FailedToGetPerson
	}

	err = n.personRepo.ClearPersonDisability(personId)
	if err != nil {
		return utils.FailedToClearDisability
	}

	for _, disabilityId := range disabilities {
		disability := model.Disability{
			Id: disabilityId,
		}

		err = n.personRepo.UpsertPersonDisability(disability, personId)
		if err != nil {
			return utils.FailedToUpsertDisability
		}
	}

	err = n.personRepo.UpdatePerson(person, personId)
	if err != nil {
		return utils.FailedToUpdatePerson
	}

	return utils.Error{}
}

func (n *personService) DeletePerson(personId int) utils.Error {
	person, err := n.personRepo.GetPersonById(personId)
	if err != nil {
		return utils.FailedToGetPerson
	}

	err = n.personRepo.DeletePerson(personId)
	if err != nil {
		return utils.FailedToDeletePerson
	}

	err = n.userRepo.DeleteUser(person.UserId)
	if err != nil {
		return utils.FailedToDeleteUser
	}

	err = n.addressRepo.DeleteAddress(*person.AddressId)
	if err != nil {
		return utils.FailedToDeleteAddress
	}

	return utils.Error{}
}
