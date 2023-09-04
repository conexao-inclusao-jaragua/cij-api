package main

import (
	"cij_api/src/config"
	"cij_api/src/controller"
	"cij_api/src/database"
	"cij_api/src/model"
	"cij_api/src/repo"
	"cij_api/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load enviroment variables", err)
	}

	db := database.ConnectionDB(&loadConfig)

	db.AutoMigrate(&model.User{})

	startServer(db)
}

func startServer(db *gorm.DB) {
	app := fiber.New()

	userRepo := repo.NewUserRepo(db)
	userController := controller.NewUserController(userRepo)

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	routes := router.NewRouter(app, userController)

	err := routes.Listen(":3040")
	if err != nil {
		panic(err)
	}
}
