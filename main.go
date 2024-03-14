package main

import (
	"example.com/todo/db"
	"example.com/todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize DB
	db.InitDB()

	// Create a server instance
	server := gin.Default()

	// Setup routes
	routes.RegisterRoutes(server)

	// Start listening for requests
	server.Run(":8080")

}
