package router

import (
	"cij_api/src/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, userController *controller.UserController) *fiber.App {
	router.Get("/", userController.CreateUser)

	// router.Post("/timeline", timelineController.CreateTimeline)
	// router.Put("/timeline/:id", timelineController.UpdateTimeline)
	// router.Get("/timeline", timelineController.ListTimelines)
	// router.Get("/timeline/:id", timelineController.GetTimeline)
	// router.Delete("/timeline/:id", timelineController.DeleteTimeline)

	return router
}
