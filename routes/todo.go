package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/todo/models"
	"github.com/gin-gonic/gin"
)

func createToDo(context *gin.Context) {
	var toDo models.ToDo
	// get the data from the context body and put it in todo variable
	err := context.ShouldBindJSON(&toDo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	// Save the todo item in db
	err = toDo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create todo item", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "todo item created successfully", "todo": toDo})
}

func getTodos(context *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todos", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetched todos successfully", "todos": todos})
}

func getTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "fetched todo successfully", "todo": todo})
}

func updateTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println("Id of the todo to update: ", todoId)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	todo, err := models.GetTodoById(todoId)
	// fmt.Println("Todo before update: ", todo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo", "error": err.Error()})
		return
	}
	var updatedTodo models.ToDo
	err = context.ShouldBindJSON(&updatedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	// fmt.Println(updatedTodo)
	updatedTodo.ID = todo.ID
	// fmt.Println(updatedTodo)
	err = updatedTodo.Update()
	// fmt.Println(updatedTodo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated todo successfully", "todo": updatedTodo})
}

func deleteTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println("Id of the todo to delete: ", todoId)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo", "error": err.Error()})
		return
	}
	err = todo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted todo successfully"})
}
