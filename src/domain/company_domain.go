package domain

import "cij_api/src/model"

type CompanyRepo interface {
	CreateCompany(createCompany model.Company) error
	ListCompanies() ([]model.Company, error)
	GetCompanyByEmail(email string) (model.Company, error)
}

type CompanyService interface {
	CreateCompany(createCompany model.Company) error
	ListCompanies() ([]model.Company, error)
	GetCompanyByEmail(email string) (model.Company, error)
}
