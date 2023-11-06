package service

import (
	"cij_api/src/auth"
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

func NewCompanyService(companyRepo domain.CompanyRepo) domain.CompanyService {
	return &companyService{
		companyRepo: companyRepo,
	}
}

type companyService struct {
	companyRepo domain.CompanyRepo
}

func (s *companyService) ListCompanies() ([]model.Company, error) {
	companies, err := s.companyRepo.ListCompanies()
	if err != nil {
		return companies, errors.New("failed to list companies")
	}

	return companies, nil
}

func (n *companyService) CreateCompany(createCompany model.Company) error {
	hashedPassword, err := auth.EncryptPassword(createCompany.Password)
	if err != nil {
		return errors.New("error on encrypt company password")
	}

	createCompany.Password = hashedPassword

	err = n.companyRepo.CreateCompany(createCompany)

	if err != nil {
		return errors.New("error on create company")
	}

	return nil
}

func (n *companyService) GetCompanyByEmail(email string) (model.Company, error) {
	company, err := n.companyRepo.GetCompanyByEmail(email)
	if err != nil {
		return company, errors.New("failed to get company by email")
	}

	return company, nil
}
