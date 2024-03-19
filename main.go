package main

import (
	"fmt"
	"os"

	"example.com/todo/db"
	"example.com/todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Get port address from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("PORT environment variable not set!, Defaulting to 8080")
	}

	// Initialize DB
	db.InitDB()

	// Create a server instance
	server := gin.Default()

	// Setup routes
	routes.RegisterRoutes(server)

	// Start listening for requests
	fmt.Println("Listening on port: ", port)
	server.Run(":" + port)
}
