package controller

import (
	"cij_api/src/model"
	"cij_api/src/service"
	"cij_api/src/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type PersonController struct {
	personService service.PersonService
}

func NewPersonController(personService service.PersonService) *PersonController {
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
// @Success 200 {object} MessageResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} MessageResponse
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

	if err := validatePersonRequiredFields(personRequest); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.validatePerson(personRequest); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := validateAddress(personRequest.Address); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.validatePersonDisabilities(personRequest.Disabilities); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.CreatePerson(personRequest); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
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
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /people [get]
func (n *PersonController) ListPeople(ctx *fiber.Ctx) error {
	var response model.Response

	people, err := n.personService.ListPeople()
	if err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
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
// @Success 200 {object} MessageResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} MessageResponse
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

	person, errPerson := n.personService.GetPersonById(idInt)
	if errPerson.Code != "" {
		response = model.Response{
			Message: errPerson.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if person.Id == 0 {
		response = model.Response{
			Message: "person not found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	if err := n.validatePerson(personRequest); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.UpdatePerson(personRequest, idInt); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
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
// @Success 200 {object} MessageResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} MessageResponse
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

	person, errPerson := n.personService.GetPersonById(idInt)
	if errPerson.Code != "" {
		response = model.Response{
			Message: errPerson.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if person.Id == 0 {
		response = model.Response{
			Message: "person not found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	if err := validateAddress(addressRequest); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.UpdatePersonAddress(addressRequest, idInt); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
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
// @Param disabilities body []model.DisabilityRequest true "Disabilities"
// @Param id path string true "Person ID"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} MessageResponse
// @Router /people/:id/disabilities [put]
func (n *PersonController) UpdatePersonDisabilities(ctx *fiber.Ctx) error {
	var disabilities []model.DisabilityRequest
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

	person, errPerson := n.personService.GetPersonById(idInt)
	if errPerson.Code != "" {
		response = model.Response{
			Message: errPerson.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if person.Id == 0 {
		response = model.Response{
			Message: "person not found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	if err := n.validatePersonDisabilities(disabilities); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.personService.UpdatePersonDisabilities(disabilities, idInt); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
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
// @Success 200 {object} MessageResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} MessageResponse
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

	person, errPerson := n.personService.GetPersonById(idInt)
	if errPerson.Code != "" {
		response = model.Response{
			Message: errPerson.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if person.Id == 0 {
		response = model.Response{
			Message: "person not found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	if err := n.personService.DeletePerson(idInt); err.Code != "" {
		response = model.Response{
			Message: err.Error(),
			Code:    err.GetCode(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		Message: "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func validatePersonRequiredFields(personRequest model.PersonRequest) utils.Error {
	if personRequest.Name == "" {
		return utils.NewError("name is required", "ERR-0001")
	}

	if personRequest.Cpf == "" {
		return utils.NewError("cpf is required", "ERR-0002")
	}

	if personRequest.Phone == "" {
		return utils.NewError("phone is required", "ERR-0003")
	}

	if personRequest.Gender == "" {
		return utils.NewError("gender is required", "ERR-0004")
	}

	if personRequest.User.Email == "" {
		return utils.NewError("email is required", "ERR-0005")
	}

	if personRequest.User.Password == "" {
		return utils.NewError("password is required", "ERR-0006")
	}

	return utils.Error{}
}

func (c *PersonController) validatePerson(personRequest model.PersonRequest) utils.Error {
	if len(personRequest.Cpf) != 11 {
		return utils.NewError("cpf must have 11 digits", "ERR-0007")
	}

	person, err := c.personService.GetPersonByCpf(personRequest.Cpf)
	if err.Code != "" {
		return utils.FailedToGetPerson
	}

	if person.Id != 0 {
		return utils.NewError("cpf already registered", "ERR-0008")
	}

	if len(personRequest.Phone) != 13 {
		return utils.NewError("phone must have 13 digits", "ERR-0009")
	}

	if !personRequest.Gender.IsValid() {
		return utils.NewError("gender is not valid", "ERR-0010")
	}

	user, err := c.personService.GetUserByEmail(personRequest.User.Email)
	if err.Code != "" {
		return utils.NewError("failed to get user by email", "ERR-0011")
	}

	if user.Id != 0 {
		return utils.NewError("email already registered", "ERR-0012")
	}

	return utils.Error{}
}

func validateAddress(addressRequest model.AddressRequest) utils.Error {
	if addressRequest.Street == "" {
		return utils.NewError("street is required", "ERR-0013")
	}

	if addressRequest.Number == "" {
		return utils.NewError("number is required", "ERR-0014")
	}

	if addressRequest.Neighborhood == "" {
		return utils.NewError("neighborhood is required", "ERR-0015")
	}

	if addressRequest.City == "" {
		return utils.NewError("city is required", "ERR-0016")
	}

	if addressRequest.State == "" {
		return utils.NewError("state is required", "ERR-0017")
	}

	if addressRequest.ZipCode == "" {
		return utils.NewError("zip code is required", "ERR-0018")
	}

	if len(addressRequest.ZipCode) != 8 {
		return utils.NewError("zip code must have 8 digits", "ERR-0019")
	}

	if len(addressRequest.State) != 2 {
		return utils.NewError("state must have 2 digits", "ERR-0020")
	}

	return utils.Error{}
}

func (n *PersonController) validatePersonDisabilities(disabiliesRequest []model.DisabilityRequest) utils.Error {
	for _, disability := range disabiliesRequest {
		disability, err := n.personService.GetDisabilityById(disability.Id)
		if err.Code != "" {
			return utils.NewError("failed to get disability", "ERR-0021")
		}

		if disability.Id == 0 {
			return utils.NewError("disability not found", "ERR-0022")
		}
	}

	return utils.Error{}
}
