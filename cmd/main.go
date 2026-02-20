package main

import (
	"fmt"
	"os"

	//.env file support
	"github.com/joho/godotenv"

	//gin http support
	"github.com/gin-gonic/gin"

	"narrativeboard/internal/routes"
)

func main() {
	godotenv.Load()

	fmt.Println("Hello World!")
	router := gin.Default()

	routes.RegisterRoutes(router)
	port := os.Getenv("PORT")

	if err := router.Run(fmt.Sprintf(":%v", port)); err != nil {
		fmt.Errorf("Error while running server: %v", err)
	}
}
