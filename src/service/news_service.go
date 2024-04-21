package service

import (
	"cij_api/src/model"
	"cij_api/src/repo"
	"cij_api/src/utils"
	"errors"

	"mime/multipart"
	"strings"
)

type NewsService interface {
	ListNews() ([]model.NewsResponse, utils.Error)
	CreateNews(model.NewsRequest, map[string]multipart.FileHeader) utils.Error
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
	filesService := NewFilesService()
	news := createNews.ToModel()

	for fileName, image := range images {
		openedFile, err := image.Open()
		if err != nil {
			return errors.New("failed to open file")
		}

		defer openedFile.Close()

		switch fileName {
		case "banner":
			fileUrl, err := filesService.UploadFile(openedFile, "cij/news/banner/"+strings.Split(image.Filename, ".")[0])
			if err != nil {
				return errors.New("failed to upload file")
			}

			news.Banner = fileUrl

		case "author_image":
			fileUrl, err := filesService.UploadFile(openedFile, "cij/news/author_image/"+strings.Split(image.Filename, ".")[0])
			if err != nil {
				return errors.New("failed to upload file")
			}

			news.AuthorImage = fileUrl

		default:
			return errors.New("invalid file name")
		}
	}

	err := n.newsRepo.CreateNews(news)
	if err.Code != "" {
		return err
	}

	return utils.Error{}
}
