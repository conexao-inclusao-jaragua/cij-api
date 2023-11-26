package domain

import "cij_api/src/model"

type CompanyRepo interface {
	CreateCompany(createCompany model.Company) error
	ListCompanies() ([]model.Company, error)
}

type CompanyService interface {
	CreateCompany(createCompany model.CompanyRequest) error
	ListCompanies() ([]model.CompanyResponse, error)
}
