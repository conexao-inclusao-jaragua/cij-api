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

func (n *newsService) ListNews() ([]model.NewsResponse, error) {
	newsResponse := []model.NewsResponse{}

	news, err := n.newsRepo.ListNews()
	if err != nil {
		return newsResponse, errors.New("failed to list news")
	}

	for _, news := range news {
		newsResponse = append(newsResponse, news.ToResponse())
	}

	return newsResponse, nil
}
