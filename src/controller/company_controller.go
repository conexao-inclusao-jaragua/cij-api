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
