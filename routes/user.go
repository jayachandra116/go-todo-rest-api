package routes

import (
	"net/http"
	"strconv"

	"example.com/todo/models"
	"example.com/todo/utils"
	"github.com/gin-gonic/gin"
)

func createNewUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request body", "error": err.Error()})
		return
	}
	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not hash user password", "error": err.Error()})
		return
	}
	user.Password = hashedPassword
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Created user successfully", "user": user})
}

func getUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id", "error": err.Error()})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user fetched successfully", "user": user})
}

func getUserByEmail(context *gin.Context) {
	email := context.Param("email")
	user, err := models.GetUserByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user by email", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user fetched successfully", "user": user})
}

func updateUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	outdatedUser, err := models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch user to update", "error": err.Error()})
		return
	}
	var updatedUser models.User
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	updatedUser.ID = outdatedUser.ID
	updatedUser.Email = user.Email
	updatedUser.Password = user.Password
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated user successfully"})
}

func deleteUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	var user *models.User
	user, err = models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user", "error": err.Error()})
		return
	}
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted user successfully"})
}

func changeUserPassword(context *gin.Context) {
	type NewPassword struct {
		Value string
	}
	var newPassword NewPassword
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	// fmt.Println("User id to change password: ", id)
	err = context.ShouldBindJSON(&newPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	// fmt.Println("New password: ", newPassword)
	err = models.ChangePassword(id, newPassword.Value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not change password of user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Changed password successfully"})
}

func loginUser(context *gin.Context) {
	var err error
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse login request data", "error": err.Error()})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
