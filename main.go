package main

import (
	"cij_api/src/config"
	"cij_api/src/database"
	"cij_api/src/model"
	"cij_api/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

const PRD_ENV = "prd"

var env string

// @title GO Clean API
// @version 1.0
// @description This the API for the CIJ project
// @contact.name API Support
// @contact.email cauakathdev@gmail.com
// @host localhost:3040
// @BasePath /
func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load enviroment variables", err)
	}

	db := database.ConnectionDB(&loadConfig)

	migrateDb(db)

	startServer(db)
}

func migrateDb(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Address{})
	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.Company{})
	db.AutoMigrate(&model.News{})
	db.AutoMigrate(&model.Role{})

	createDefaultRoles(db)
}

func createDefaultRoles(db *gorm.DB) {
	db.Create(&model.Role{
		Name: "person",
	})

	db.Create(&model.Role{
		Name: "company",
	})

	db.Create(&model.Role{
		Name: "admin",
	})
}

func startServer(db *gorm.DB) {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	routes := router.NewRouter(app, db)

	if env == PRD_ENV {
		certPath := "../certs/server.pem"
		keyPath := "../certs/server.key"

		err := routes.ListenTLS(":3040", certPath, keyPath)
		if err != nil {
			panic(err)
		}

		return
	}

	err := routes.Listen(":3040")
	if err != nil {
		panic(err)
	}
}
