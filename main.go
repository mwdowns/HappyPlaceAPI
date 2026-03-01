package main

import (
	"fmt"
	"happyplace/api/db"
	"happyplace/api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to your HappyPlace!")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	if os.Getenv("APP_ENV") == "develop" {
		router.Run(os.Getenv("APP_PORT"))
	}
}
