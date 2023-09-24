package auth

import (
	"cij_api/src/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Authenticate(ctx *fiber.Ctx) error {
	var userCredentials model.Credentials
	var response model.LoginResponse

	if err := ctx.BodyParser(&userCredentials); err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := c.authService.Authenticate(userCredentials)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	token, err := c.authService.GenerateToken(user)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	response = model.LoginResponse{
		Token:    token,
		UserInfo: user,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
