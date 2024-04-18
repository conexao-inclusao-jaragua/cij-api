package repo

import (
	"cij_api/src/domain"
	"cij_api/src/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type personDisabilityRepo struct {
	db *gorm.DB
}

func NewPersonDisabilityRepo(db *gorm.DB) domain.PersonDisabilityRepo {
	return &personDisabilityRepo{
		db: db,
	}
}

func (n *personDisabilityRepo) GetPersonDisabilities(personId int) ([]model.PersonDisability, error) {
	var disabilities []model.PersonDisability

	err := n.db.Model(model.PersonDisability{}).Preload("Disability").Where("person_id = ?", personId).Find(&disabilities).Error
	if err != nil {
		return disabilities, err
	}

	return disabilities, nil
}

func (n *personDisabilityRepo) GetDisabilityById(disabilityId int) (model.Disability, error) {
	var disability model.Disability

	err := n.db.Model(model.Disability{}).Where("id = ?", disabilityId).Find(&disability).Error
	if err != nil {
		return disability, err
	}

	return disability, nil
}

func (n *personDisabilityRepo) UpsertPersonDisability(personDisability model.PersonDisability) error {
	n.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "person_id"}, {Name: "disability_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"acquired"}),
	}).Create(&personDisability)

	return nil
}

func (n *personDisabilityRepo) ClearPersonDisability(personId int) error {
	if err := n.db.Where("person_id = ?", personId).Delete(&model.PersonDisability{}).Error; err != nil {
		return err
	}

	return nil
}
