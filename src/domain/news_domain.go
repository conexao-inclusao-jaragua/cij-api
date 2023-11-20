package domain

import "cij_api/src/model"

type NewsRepo interface {
	ListNews() ([]model.News, error)
}

type NewsService interface {
	ListNews() ([]model.News, error)
}
