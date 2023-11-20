package controller

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct {
	newsService domain.NewsService
}

func NewNewsController(newsService domain.NewsService) *NewsController {
	return &NewsController{
		newsService: newsService,
	}
}

func (n *NewsController) ListNews(ctx *fiber.Ctx) error {
	var response model.Response

	allNews, err := n.newsService.ListNews()
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if len(allNews) == 0 {
		response = model.Response{
			Message: "No news were found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	newsResponse := []model.NewsResponse{}

	for _, news := range allNews {
		newsResponse = append(newsResponse, news.ToResponse())
	}

	response = model.Response{
		Message: "success",
		Data:    newsResponse,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
