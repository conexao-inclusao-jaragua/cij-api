package controller

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PersonController struct {
	personService domain.PersonService
}

func NewPersonController(personService domain.PersonService) *PersonController {
	return &PersonController{
		personService: personService,
	}
}

func (n *PersonController) CreatePerson(ctx *fiber.Ctx) error {
	var personRequest model.PersonRequest
	var response model.Response

	if err := ctx.BodyParser(&personRequest); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.CreatePerson(personRequest); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		Message: "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *PersonController) ListPeople(ctx *fiber.Ctx) error {
	var response model.Response

	people, err := n.personService.ListPeople()
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if len(people) == 0 {
		response = model.Response{
			Message: "no people were found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	response = model.Response{
		Message: "success",
		Data:    people,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *PersonController) UpdatePerson(ctx *fiber.Ctx) error {
	var personRequest model.PersonRequest
	var response model.Response

	if err := ctx.BodyParser(&personRequest); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	personId := ctx.Params("id")

	idInt, err := strconv.Atoi(personId)
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.UpdatePerson(personRequest, idInt); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		Message: "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *PersonController) DeletePerson(ctx *fiber.Ctx) error {
	var response model.Response

	personId := ctx.Params("id")

	idInt, err := strconv.Atoi(personId)
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.DeletePerson(idInt); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		Message: "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
