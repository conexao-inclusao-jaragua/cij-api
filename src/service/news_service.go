package service

import (
	"cij_api/src/model"
	"cij_api/src/repo"
	"errors"
)

type NewsService interface {
	ListNews() ([]model.NewsResponse, error)
}

type newsService struct {
	newsRepo repo.NewsRepo
}

func NewNewsService(newsRepo repo.NewsRepo) NewsService {
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
