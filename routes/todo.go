package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"example.com/todo/models"
	"example.com/todo/utils"
	"github.com/gin-gonic/gin"
)

// Creates a new todo item
func createToDo(context *gin.Context) {
	fmt.Println("Creating a new todo item ...")
	// Get the token from authorization header from request
	fmt.Println("Extracting the token from request...")
	bearerToken := context.Request.Header.Get("Authorization")
	l := strings.Split(bearerToken, " ")
	token := l[1]
	fmt.Println("Got bearer token: ", token)
	// Verify token is valid
	fmt.Println("Verifying the token ...")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token", "error": err.Error()})
		return
	}
	fmt.Println("Got user ID from the token: ", userId)
	// if valid, set the userId from the token to the todo item
	var todo models.ToDo
	todo.UserID = userId
	type request struct {
		TodoContent string `json:"todoContent" binding:"required"`
	}
	// get the data from the context body and put it in todo variable
	fmt.Println("Parsing the request body ...")
	var requestBody request
	err = context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	// Save the todo item in db
	todo.Content = requestBody.TodoContent
	fmt.Println("Saved todo item ...")
	fmt.Println(todo)
	err = todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create todo item", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "todo item created successfully", "todo": todo})
}

func getTodoById(context *gin.Context) {
	fmt.Println("Getting todo item by id ...")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	fmt.Println("Got todo id: ", todoId)
	// Get the token from authorization header from request
	bearerToken := context.Request.Header.Get("Authorization")
	l := strings.Split(bearerToken, " ")
	token := l[1]
	fmt.Println("Got bearer token: ", token)
	// Verify token is valid
	fmt.Println("Verifying token ...")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token data", "error": err.Error()})
		return
	}
	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo", "error": err.Error()})
		return
	}
	if !(todo.UserID == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "fetched todo successfully", "todo": todo})
}

// Update todo item
func updateTodoById(context *gin.Context) {
	fmt.Println("Updating todo ...")
	// Parse request parameters -> todo id
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	fmt.Println("Id of the todo to update: ", todoId)
	// Verify Token
	// Get the token from authorization header from request
	bearerToken := context.Request.Header.Get("Authorization")
	l := strings.Split(bearerToken, " ")
	token := l[1]
	fmt.Println("Got bearer token: ", token)
	// Verify token is valid
	fmt.Println("Verifying token ...")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token data", "error": err.Error()})
		return
	}
	todo, err := models.GetTodoById(todoId)
	// fmt.Println("Todo before update: ", todo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo", "error": err.Error()})
		return
	}
	if !(todo.UserID == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// var updatedTodo models.ToDo
	fmt.Println("Parsing request body ...")
	type request struct {
		Content string `json:"content"`
	}
	var r request
	err = context.ShouldBindJSON(&r)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	todo.Content = r.Content
	fmt.Println("Updating todo ...")
	err = todo.Update()
	// fmt.Println(updatedTodo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated todo successfully", "todo": todo})
}

func deleteTodoById(context *gin.Context) {
	fmt.Println("Deleting todo ...")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	fmt.Println("Id of the todo to delete: ", todoId)
	// Verify Token
	// Get the token from authorization header from request
	bearerToken := context.Request.Header.Get("Authorization")
	l := strings.Split(bearerToken, " ")
	token := l[1]
	fmt.Println("Got bearer token: ", token)
	// Verify token is valid
	fmt.Println("Verifying token ...")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token data", "error": err.Error()})
		return
	}
	// Is user authorized?
	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo", "error": err.Error()})
		return
	}
	if !(todo.UserID == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err = todo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted todo successfully"})
}
