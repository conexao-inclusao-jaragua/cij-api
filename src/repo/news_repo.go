package repo

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"

	"gorm.io/gorm"
)

type newsRepo struct {
	db *gorm.DB
}

func NewNewsRepo(db *gorm.DB) domain.NewsRepo {
	return &newsRepo{
		db: db,
	}
}

func (r *newsRepo) ListNews() ([]model.News, error) {
	var news []model.News

	err := r.db.Model(model.News{}).Find(&news).Error
	if err != nil {
		return news, errors.New("error on list news from database")
	}

	return news, nil
}
