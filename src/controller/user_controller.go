package controller

import (
	"cij_api/src/domain"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userRepo domain.UserRepo
}

func NewUserController(userRepo domain.UserRepo) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func (n *UserController) CreateUser(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("Hello World")
}
