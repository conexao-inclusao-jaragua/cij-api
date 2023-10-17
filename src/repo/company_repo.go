package repo

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) domain.CompanyRepo {
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

	err := n.db.Model(model.Company{}).Find(&companies).Error
	if err != nil {
		return companies, errors.New("error on list companies from database")
	}

	return companies, nil
}

func (n *companyRepo) GetCompanyByEmail(email string) (model.Company, error) {
	var company model.Company

	err := n.db.Model(model.Company{}).Where("email = ?", email).Find(&company).Error
	if err != nil {
		return company, errors.New("failed to get the company")
	}

	return company, nil
}
