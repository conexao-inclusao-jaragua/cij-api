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
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	companyRepo := repo.NewCompanyRepo(db)
	companyService := service.NewCompanyService(companyRepo)
	companyController := controller.NewCompanyController(companyService)

	newsRepo := repo.NewNewsRepo(db)
	newsService := service.NewNewsService(newsRepo)
	newsController := controller.NewNewsController(newsService)

	authService := auth.NewAuthService(userRepo, companyRepo)
	authController := auth.NewAuthController(*authService)

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Server running",
		})
	})

	router.Post("/login/:role", authController.Authenticate)
	router.Post("/get-user-data", authController.GetUserData)

	api := router.Group("/users")
	{
		api.Post("/create", userController.CreateUser)

		api.Use(middleware.AuthUser)
		api.Get("/list", userController.ListUsers)
	}

	api = router.Group("/companies")
	{
		api.Post("/create", companyController.CreateCompany)

		api.Use(middleware.AuthCompany)
		api.Get("/list", companyController.ListCompanies)
	}

	api = router.Group("news")
	{
		api.Get("/list", newsController.ListNews)
	}

	return router
}
