package auth

import (
	"cij_api/src/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService AuthService
}

type TokenRequest struct {
	Token string `json:"token"`
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Authenticate(ctx *fiber.Ctx) error {
	var credentials model.Credentials
	var response model.LoginResponse
	var role = ctx.Params("role")

	if role != "user" && role != "company" {
		response = model.LoginResponse{
			Message: "role not found",
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := ctx.BodyParser(&credentials); err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	user, company, err := c.authService.Authenticate(credentials, role)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	token, err := c.authService.GenerateToken(role, user, company)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if user.Email != "" {
		response = model.LoginResponse{
			Token:    token,
			UserInfo: user.ToResponse(),
		}
	} else if company.Email != "" {
		response = model.LoginResponse{
			Token:    token,
			UserInfo: company.ToResponse(),
		}
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (c *AuthController) GetUserData(ctx *fiber.Ctx) error {
	var token TokenRequest
	var response model.LoginResponse

	if err := ctx.BodyParser(&token); err != nil {
		response = model.LoginResponse{
			Message: "token not found",
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	user, company, err := c.authService.GetUserData(token.Token)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if user.Email != "" {
		response = model.LoginResponse{
			UserInfo: user.ToResponse(),
		}
	} else if company.Email != "" {
		response = model.LoginResponse{
			UserInfo: company.ToResponse(),
		}
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
