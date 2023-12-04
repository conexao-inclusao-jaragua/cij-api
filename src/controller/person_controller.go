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

// CreatePerson
// @Summary Create a new person.
// @Description create a new person and their user.
// @Tags People
// @Accept */*
// @Produce json
// @Param person body model.PersonRequest true "Person"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /people [post]
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

// ListPeople
// @Summary List all registered people.
// @Description list all registered people and their users.
// @Tags People
// @Accept */*
// @Produce json
// @Success 200 {array} model.PersonResponse
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "internal server error"
// @Router /people [get]
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

// UpdatePerson
// @Summary Update a person.
// @Description update an existent person and their user.
// @Tags People
// @Accept */*
// @Produce json
// @Param person body model.PersonRequest true "Person"
// @Param id path string true "Person ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /people/:id [put]
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

// UpdatePersonAddress
// @Summary Update a person address.
// @Description update an existent person address.
// @Tags People
// @Accept */*
// @Produce json
// @Param address body model.AddressRequest true "Address"
// @Param id path string true "Person ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /people/:id/address [put]
func (n *PersonController) UpdatePersonAddress(ctx *fiber.Ctx) error {
	var addressRequest model.AddressRequest
	var response model.Response

	if err := ctx.BodyParser(&addressRequest); err != nil {
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

	if err := n.personService.UpdatePersonAddress(addressRequest, idInt); err != nil {
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

// UpdatePersonDisabilities
// @Summary Update a person disabilities.
// @Description update an existent person disabilities.
// @Tags People
// @Accept */*
// @Produce json
// @Param disabilities body []int true "Disabilities"
// @Param id path string true "Person ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /people/:id/disabilities [put]
func (n *PersonController) UpdatePersonDisabilities(ctx *fiber.Ctx) error {
	var disabilities []int
	var response model.Response

	if err := ctx.BodyParser(&disabilities); err != nil {
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

	if err := n.personService.UpdatePersonDisabilities(disabilities, idInt); err != nil {
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

// DeletePerson
// @Summary Delete a person.
// @Description delete an existent person and their user.
// @Tags People
// @Accept */*
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /people/:id [delete]
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
