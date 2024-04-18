package repo

import (
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type CompanyRepo interface {
	CreateCompany(createCompany model.Company) error
	ListCompanies() ([]model.Company, error)
	GetCompanyById(companyId int) (model.Company, error)
	GetCompanyByUserId(userId int) (model.Company, error)
	UpdateCompany(company model.Company, companyId int) error
	DeleteCompany(companyId int) error
}

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) CompanyRepo {
	return &companyRepo{
		db: db,
	}
}

func (n *companyRepo) CreateCompany(createCompany model.Company) error {
	if err := n.db.Create(&createCompany).Error; err != nil {
		return errors.New("failed to create company")
	}

	return nil
}

func (n *companyRepo) ListCompanies() ([]model.Company, error) {
	var companies []model.Company

	err := n.db.Model(model.Company{}).Preload("User").Preload("Address").Find(&companies).Error
	if err != nil {
		return companies, errors.New("error on list companies from database")
	}

	return companies, nil
}

func (n *companyRepo) GetCompanyById(companyId int) (model.Company, error) {
	var company model.Company

	err := n.db.Model(model.Company{}).Preload("User").Where("id = ?", companyId).Find(&company).Error
	if err != nil {
		return company, errors.New("failed to get the company")
	}

	return company, nil
}

func (n *companyRepo) GetCompanyByUserId(userId int) (model.Company, error) {
	var company model.Company

	err := n.db.Model(model.Company{}).Preload("User").Where("user_id = ?", userId).Find(&company).Error
	if err != nil {
		return company, errors.New("failed to get the company")
	}

	return company, nil
}

func (n *companyRepo) UpdateCompany(company model.Company, companyId int) error {
	if err := n.db.Model(model.Company{}).Where("id = ?", companyId).Updates(company).Error; err != nil {
		return errors.New("failed to update the company")
	}

	return nil
}

func (n *companyRepo) DeleteCompany(companyId int) error {
	if err := n.db.Model(model.Company{}).Where("id = ?", companyId).Delete(&model.Company{}).Error; err != nil {
		return errors.New("failed to delete the company")
	}

	return nil
}
