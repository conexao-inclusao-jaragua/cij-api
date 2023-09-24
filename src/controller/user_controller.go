package controller

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService domain.UserService
}

func NewUserController(userService domain.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (n *UserController) CreateUser(ctx *fiber.Ctx) error {
	var userRequest model.User
	var response model.Response

	if err := ctx.BodyParser(&userRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.userService.CreateUser(userRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusCreated,
		Message:    "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *UserController) ListUsers(ctx *fiber.Ctx) error {
	var response model.Response

	users, err := n.userService.ListUsers()
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       users,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
