package controller

import (
	"cij_api/src/model"
	"cij_api/src/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct {
	newsService service.NewsService
}

func NewNewsController(newsService service.NewsService) *NewsController {
	return &NewsController{
		newsService: newsService,
	}
}

// ListNews
// @Summary List all registered news.
// @Description list all registered news.
// @Tags News
// @Accept application/json
// @Produce json
// @Success 200 {array} model.NewsResponse
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "internal server error"
// @Router /news [get]
func (n *NewsController) ListNews(ctx *fiber.Ctx) error {
	var response model.Response

	news, err := n.newsService.ListNews()
	if err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.Code,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		Message: "success",
		Data:    news,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *NewsController) CreateNews(ctx *fiber.Ctx) error {
	var request model.NewsRequest
	var response model.Response

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(model.Response{
			Message: err.Error(),
		})
	}

	err := n.newsService.CreateNews(request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(model.Response{
			Message: err.Error(),
		})
	}

	response = model.Response{
		Message: "success",
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}