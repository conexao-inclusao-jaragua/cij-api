package router

import (
	"cij_api/src/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, userController *controller.UserController) *fiber.App {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("Hello world")
	})

	router.Post("/user", userController.CreateUser)

	return router
}
