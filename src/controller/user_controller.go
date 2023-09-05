package controller

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	var userRequest model.User
	var response model.Response

	if err := ctx.BodyParser(&userRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	hashedPassword, err := encryptPassword(userRequest.Password)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	userRequest.Password = hashedPassword

	if err := n.userRepo.CreateUser(userRequest); err != nil {
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

func encryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error on encrypt password")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		return "", errors.New("error on encrypt password")
	}

	return string(hashedPassword), nil
}
