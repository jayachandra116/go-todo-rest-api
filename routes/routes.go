package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Create a new todo item
	server.POST("/todos", createToDo)

	// Get all the to-do items
	server.GET("/todos", getTodos)

	// Get a single todo item
	server.GET("/todos/:id", getTodo)

	// Update a single todo item
	server.PUT("/todos/:id", updateTodo)

	// Delete a single todo item
	server.DELETE("/todos/:id", deleteTodo)
}
