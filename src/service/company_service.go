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
}

func NewCompanyService(companyRepo domain.CompanyRepo, userRepo domain.UserRepo) domain.CompanyService {
	return &companyService{
		companyRepo: companyRepo,
		userRepo:    userRepo,
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

		companiesResponse = append(companiesResponse, company.ToResponse(user))
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

	createCompany.User.Id = userId

	err = n.companyRepo.CreateCompany(createCompany.ToCompany())
	if err != nil {
		n.userRepo.DeleteUser(userId)

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
