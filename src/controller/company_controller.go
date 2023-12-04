package controller

import (
	"cij_api/src/domain"
	"cij_api/src/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	companyService domain.CompanyService
}

func NewCompanyController(companyService domain.CompanyService) *CompanyController {
	return &CompanyController{
		companyService: companyService,
	}
}

// CreateCompany
// @Summary Create a new company.
// @Description create a new company and their user.
// @Tags Companies
// @Accept */*
// @Produce json
// @Param company body model.CompanyRequest true "Company"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /companies [post]
func (n *CompanyController) CreateCompany(ctx *fiber.Ctx) error {
	var companyRequest model.CompanyRequest
	var response model.Response

	if err := ctx.BodyParser(&companyRequest); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.companyService.CreateCompany(companyRequest); err != nil {
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

// ListCompanies
// @Summary List all registered companies.
// @Description list all registered companies and their users.
// @Tags Companies
// @Accept */*
// @Produce json
// @Success 200 {array} model.CompanyResponse
// @Failure 404 {object} string "not found"
// @Failure 500 {object} string "internal server error"
// @Router /companies [get]
func (n *CompanyController) ListCompanies(ctx *fiber.Ctx) error {
	var response model.Response

	companies, err := n.companyService.ListCompanies()
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if len(companies) == 0 {
		response = model.Response{
			Message: "no companies were found",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	response = model.Response{
		Message: "success",
		Data:    companies,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// UpdateCompany
// @Summary Update a company.
// @Description update an existent company and their user.
// @Tags Companies
// @Accept */*
// @Produce json
// @Param company body model.CompanyRequest true "Company"
// @Param id path string true "Company ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /companies/:id [put]
func (n *CompanyController) UpdateCompany(ctx *fiber.Ctx) error {
	var companyRequest model.CompanyRequest
	var response model.Response

	if err := ctx.BodyParser(&companyRequest); err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	companyId := ctx.Params("id")

	idInt, err := strconv.Atoi(companyId)
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.companyService.UpdateCompany(companyRequest, idInt); err != nil {
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

// DeleteCompany
// @Summary Delete a company.
// @Description delete an existent company and their user.
// @Tags Companies
// @Accept */*
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} string "success"
// @Failure 400 {object} string "bad request"
// @Failure 500 {object} string "internal server error"
// @Router /companies/:id [delete]
func (n *CompanyController) DeleteCompany(ctx *fiber.Ctx) error {
	var response model.Response

	companyId := ctx.Params("id")

	idInt, err := strconv.Atoi(companyId)
	if err != nil {
		response = model.Response{
			Message: err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.companyService.DeleteCompany(idInt); err != nil {
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
