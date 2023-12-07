package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shrisuradkar/RestAPI-MonogoDB-Golang/routes"
)

func main() {
	fmt.Println("Welcome to Rest API project using MongoDB as a Database")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.Default()
	routes.UserRoutes(router)

	router.Run(":" + port)

}
