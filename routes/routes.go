package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	registerTodoRoutes(server)
	registerUserRoutes(server)
}

func registerTodoRoutes(server *gin.Engine) {
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

func registerUserRoutes(server *gin.Engine) {
	// Create a new user
	server.POST("/users/signup", createNewUser)
	// Login existing user
	server.POST("/users/login", loginUser)
	// Get user by Id
	server.GET("/users/byId/:id", getUserById)
	// Get user by email
	server.GET("/users/byEmail/:email", getUserByEmail)
	// Update user
	server.PUT("/users/:id", updateUser)
	// Delete user
	server.DELETE("/users/:id", deleteUser)
	// Change user password
	server.PUT("/users/:id/changePwd", changeUserPassword)
}
