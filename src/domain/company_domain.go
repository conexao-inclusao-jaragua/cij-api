package domain

import "cij_api/src/model"

type CompanyRepo interface {
	CreateCompany(createCompany model.Company) error
	ListCompanies() ([]model.Company, error)
	GetCompanyById(companyId int) (model.Company, error)
	GetCompanyByUserId(userId int) (model.Company, error)
	UpdateCompany(company model.Company, companyId int) error
	DeleteCompany(companyId int) error
}

type CompanyService interface {
	CreateCompany(createCompany model.CompanyRequest) error
	ListCompanies() ([]model.CompanyResponse, error)
	GetCompanyByUserId(userId int) (model.Company, error)
	UpdateCompany(company model.CompanyRequest, companyId int) error
	DeleteCompany(companyId int) error
}
