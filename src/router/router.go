package router

import (
	"cij_api/src/auth"
	"cij_api/src/controller"
	"cij_api/src/middleware"
	"cij_api/src/repo"
	"cij_api/src/service"

	_ "cij_api/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	userRepo := repo.NewUserRepo(db)

	addressRepo := repo.NewAddressRepo(db)
	addressService := service.NewAddressService(addressRepo)

	personDisabilityRepo := repo.NewPersonDisabilityRepo(db)

	personRepo := repo.NewPersonRepo(db)
	personService := service.NewPersonService(personRepo, userRepo, addressRepo, personDisabilityRepo)
	personController := controller.NewPersonController(personService)

	companyRepo := repo.NewCompanyRepo(db)
	companyService := service.NewCompanyService(companyRepo, userRepo, addressRepo)
	companyController := controller.NewCompanyController(companyService)

	newsRepo := repo.NewNewsRepo(db)
	newsService := service.NewNewsService(newsRepo)
	newsController := controller.NewNewsController(newsService)

	authService := auth.NewAuthService(userRepo)
	authController := auth.NewAuthController(*authService, personService, companyService, addressService)

	router.Get("/health", HealthCheck)

	router.Get("/swagger/*", swagger.HandlerDefault)

	router.Post("/login", authController.Authenticate)
	router.Post("/get-user-data", authController.GetUserData)

	api := router.Group("/people")
	{
		api.Get("/", personController.ListPeople)
		api.Post("/", personController.CreatePerson)

		api.Use(middleware.AuthUser)
		api.Put("/:id", personController.UpdatePerson)
		api.Put("/:id/address", personController.UpdatePersonAddress)
		api.Put("/:id/disabilities", personController.UpdatePersonDisabilities)
		api.Delete("/:id", personController.DeletePerson)
	}

	api = router.Group("/companies")
	{
		api.Get("/", companyController.ListCompanies)

		api.Use(middleware.AuthAdmin)
		api.Post("/", companyController.CreateCompany)
		api.Put("/:id", companyController.UpdateCompany)
		api.Delete("/:id", companyController.DeleteCompany)
	}

	api = router.Group("/news")
	{
		api.Get("/", newsController.ListNews)
	}

	return router
}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
