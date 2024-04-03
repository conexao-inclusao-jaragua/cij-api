package service

import (
	"cij_api/src/model"
	"cij_api/src/repo"
	"cij_api/src/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"mime/multipart"
)

type NewsService interface {
	ListNews() ([]model.NewsResponse, utils.Error)
	CreateNews(createNews model.NewsRequest, map[string]multipart.FileHeader) utils.Error
}

type newsService struct {
	newsRepo repo.NewsRepo
}

func NewNewsService(newsRepo repo.NewsRepo) NewsService {
	return &newsService{
		newsRepo: newsRepo,
	}
}

func (n *newsService) ListNews() ([]model.NewsResponse, utils.Error) {
	newsResponse := []model.NewsResponse{}

	news, err := n.newsRepo.ListNews()
	if err.Code != "" {
		return newsResponse, err
	}

	for _, news := range news {
		newsResponse = append(newsResponse, news.ToResponse())
	}

	return newsResponse, utils.Error{}
}

func (n *newsService) CreateNews(createNews model.NewsRequest, images map[string]multipart.FileHeader) utils.Error {
	news := createNews.ToModel()

	err := n.newsRepo.CreateNews(news)
	if err.Code != "" {
		return err
	}

	return utils.Error{}
}
