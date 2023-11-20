package service

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
)

type newsService struct {
	newsRepo domain.NewsRepo
}

func NewNewsService(newsRepo domain.NewsRepo) domain.NewsService {
	return &newsService{
		newsRepo: newsRepo,
	}
}

func (n *newsService) ListNews() ([]model.News, error) {
	news, err := n.newsRepo.ListNews()

	if err != nil {
		return news, errors.New("failed to list news")
	}

	return news, nil
}
