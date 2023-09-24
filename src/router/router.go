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

	authService := auth.NewAuthService(userRepo)
	authController := auth.NewAuthController(*authService)

	api := router.Group("/users")
	{
		api.Post("/login", authController.Authenticate)
		api.Post("/create", userController.CreateUser)

		api.Use(middleware.AuthUser)
		api.Get("/list", userController.ListUsers)
	}

	return router
}
