package database

import (
	"cij_api/src/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *config.Config) *gorm.DB {
	dsn := config.DbConnection
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Print("Database connected")

	return client
}
