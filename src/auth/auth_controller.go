package auth

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService    AuthService
	personService  domain.PersonService
	companyService domain.CompanyService
}

type TokenRequest struct {
	Token string `json:"token"`
}

func NewAuthController(
	authService AuthService, personService domain.PersonService, companyService domain.CompanyService,
) *AuthController {
	return &AuthController{
		authService:   authService,
		personService: personService,
	}
}

func (c *AuthController) Authenticate(ctx *fiber.Ctx) error {
	var credentials model.Credentials
	var response model.LoginResponse

	if err := ctx.BodyParser(&credentials); err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := c.authService.Authenticate(credentials)
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
		UserInfo: user.ToResponse(),
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

	user, err := c.authService.GetUserData(token.Token)
	if err != nil {
		response = model.LoginResponse{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if user.RoleId == 2 {
		company, err := c.companyService.GetCompanyByUserId(user.Id)
		if err != nil {
			response = model.LoginResponse{
				Message: err.Error(),
			}

			return ctx.Status(http.StatusInternalServerError).JSON(response)
		}

		response = model.LoginResponse{
			UserInfo: company.ToResponse(user),
		}

		return ctx.Status(http.StatusOK).JSON(response)
	} else {
		person, err := c.personService.GetPersonByUserId(user.Id)
		if err != nil {
			response = model.LoginResponse{
				Message: err.Error(),
			}

			return ctx.Status(http.StatusInternalServerError).JSON(response)
		}

		response = model.LoginResponse{
			UserInfo: person.ToResponse(user),
		}

		return ctx.Status(http.StatusOK).JSON(response)
	}
}
