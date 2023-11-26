package router

import (
	"cij_api/src/auth"
	"cij_api/src/controller"
	"cij_api/src/middleware"
	"cij_api/src/repo"
	"cij_api/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(router *fiber.App, db *gorm.DB) *fiber.App {
	personRepo := repo.NewPersonRepo(db)
	userRepo := repo.NewUserRepo(db)
	personService := service.NewPersonService(personRepo, userRepo)
	personController := controller.NewPersonController(personService)

	companyRepo := repo.NewCompanyRepo(db)
	companyService := service.NewCompanyService(companyRepo, userRepo)
	companyController := controller.NewCompanyController(companyService)

	newsRepo := repo.NewNewsRepo(db)
	newsService := service.NewNewsService(newsRepo)
	newsController := controller.NewNewsController(newsService)

	authService := auth.NewAuthService(userRepo)
	authController := auth.NewAuthController(*authService, personService, companyService)

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Server running",
		})
	})

	router.Post("/login", authController.Authenticate)
	router.Post("/get-user-data", authController.GetUserData)

	api := router.Group("/people")
	{
		api.Post("/", personController.CreatePerson)
		api.Get("/", personController.ListPeople)
		api.Put("/:id", personController.UpdatePerson)
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
