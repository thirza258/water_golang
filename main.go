package main

import (
	"fmt"
	"log"
	"os"
	"water/config"
	"water/models"
	Routes "water/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := config.DbURL(config.BuildDBConfig())
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB.AutoMigrate(
		&models.Quality{},
	)

	APP_URL := os.Getenv("APP_URL")
	if len(APP_URL) == 0 {
		APP_URL = "http://localhost:8080"
	}

	r := Routes.SetupRouter()

	r.Run()
}
