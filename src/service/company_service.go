package service

import (
	"cij_api/src/auth"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

type companyService struct {
	companyRepo domain.CompanyRepo
	userRepo    domain.UserRepo
	addressRepo domain.AddressRepo
}

func NewCompanyService(companyRepo domain.CompanyRepo, userRepo domain.UserRepo, addressRepo domain.AddressRepo) domain.CompanyService {
	return &companyService{
		companyRepo: companyRepo,
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

func (s *companyService) ListCompanies() ([]model.CompanyResponse, error) {
	companiesResponse := []model.CompanyResponse{}

	companies, err := s.companyRepo.ListCompanies()
	if err != nil {
		return companiesResponse, errors.New("failed to list companies")
	}

	for _, company := range companies {
		user, err := s.userRepo.GetUserById(company.User.Id)
		if err != nil {
			return companiesResponse, errors.New("failed to get user")
		}

		companyResponse := company.ToResponse(user)

		address, err := s.addressRepo.GetAddressById(*company.AddressId)
		if err != nil {
			return companiesResponse, errors.New("failed to get address")
		}

		if address.Id != 0 {
			addressResponse := address.ToResponse()
			companyResponse.Address = addressResponse
		}

		companiesResponse = append(companiesResponse, companyResponse)
	}

	return companiesResponse, nil
}

func (n *companyService) CreateCompany(createCompany model.CompanyRequest) error {
	userInfo := createCompany.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return errors.New("error on encrypt company password")
	}

	userInfo.Password = hashedPassword
	userInfo.RoleId = 2 // 2 is the id of the role "company"

	userId, err := n.userRepo.CreateUser(userInfo)
	if err != nil {
		return errors.New("error on create user")
	}

	addressInfo := createCompany.ToAddress()

	addressId, err := n.addressRepo.UpsertAddress(addressInfo)
	if err != nil {
		n.userRepo.DeleteUser(userId)

		return errors.New("error on create address")
	}

	companyInfo := createCompany.ToModel(userInfo)
	companyInfo.UserId = userId
	companyInfo.AddressId = &addressId

	err = n.companyRepo.CreateCompany(companyInfo)
	if err != nil {
		n.userRepo.DeleteUser(userId)
		n.addressRepo.DeleteAddress(addressId)

		return errors.New("error on create company")
	}

	return nil
}

func (n *companyService) GetCompanyByUserId(userId int) (model.Company, error) {
	company, err := n.companyRepo.GetCompanyByUserId(userId)
	if err != nil {
		return company, errors.New("failed to get the company")
	}

	return company, nil
}

func (n *companyService) UpdateCompany(updateCompany model.CompanyRequest, companyId int) error {
	userInfo := updateCompany.ToUser()

	hashedPassword, err := auth.EncryptPassword(userInfo.Password)
	if err != nil {
		return errors.New("error on encrypt company password")
	}

	userInfo.Password = hashedPassword

	err = n.userRepo.UpdateUser(userInfo, companyId)
	if err != nil {
		return errors.New("failed to update the user")
	}

	addressInfo := updateCompany.ToAddress()

	company, err := n.companyRepo.GetCompanyById(companyId)
	if err != nil {
		return errors.New("failed to get the company")
	}

	addressInfo.Id = *company.AddressId

	addressId, err := n.addressRepo.UpsertAddress(addressInfo)
	if err != nil {
		return errors.New("error on create address")
	}

	companyInfo := updateCompany.ToModel(userInfo)
	companyInfo.AddressId = &addressId

	err = n.companyRepo.UpdateCompany(companyInfo, companyId)
	if err != nil {
		return errors.New("failed to update the company")
	}

	return nil
}

func (n *companyService) DeleteCompany(companyId int) error {
	company, err := n.companyRepo.GetCompanyById(companyId)
	if err != nil {
		return errors.New("failed to get the company")
	}

	err = n.companyRepo.DeleteCompany(companyId)
	if err != nil {
		return errors.New("failed to delete the company")
	}

	err = n.userRepo.DeleteUser(company.UserId)
	if err != nil {
		return errors.New("failed to delete user")
	}

	err = n.addressRepo.DeleteAddress(*company.AddressId)
	if err != nil {
		return errors.New("failed to delete address")
	}

	return nil
}
