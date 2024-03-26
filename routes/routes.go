package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Register all the routes on the given server instance
func RegisterRoutes(server *gin.Engine) {
	fmt.Println("Registering routes to the server...")
	registerUserRoutes(server)
	registerTodoRoutes(server)
}

// Register the user routes on the given server instance
func registerUserRoutes(server *gin.Engine) {
	fmt.Println("Registering user routes ...")
	// Create a new user
	server.POST("/users/signup", createNewUser)
	// Login existing user
	server.POST("/users/login", loginUser)
	// Get user by Id
	server.GET("/users/byId/:id", getUserById)
	// Get user by email
	server.GET("/users/byEmail/:email", getUserByEmail)
	// Update user
	server.PUT("/users/:id", updateUserById)
	// Delete user
	server.DELETE("/users/:id", deleteUserById)
	// Change user password
	server.PUT("/users/:id/changePwd", changeUserPassword)
}

// Register the todo routes on the given server instance
func registerTodoRoutes(server *gin.Engine) {
	fmt.Println("Registering todo routes...")
	// Create a new todo item
	server.POST("/todos", createToDo)
	// Get a single todo item
	server.GET("/todos/:id", getTodoById)
	// Update a single todo item
	server.PUT("/todos/:id", updateTodoById)
	// Delete a single todo item
	server.DELETE("/todos/:id", deleteTodoById)
}
