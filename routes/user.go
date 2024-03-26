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

func createNewUser(context *gin.Context) {
	fmt.Println("Creating new user...")
	type request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var r request
	fmt.Println("Parsing request body ...")
	err := context.ShouldBindJSON(&r)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request body", "error": err.Error()})
		return
	}
	fmt.Println("Request body:", r)
	fmt.Println("Validating request ...")
	// Validations
	if !utils.IsValidEmail(r.Email) {
		fmt.Println("email:", r.Email)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse email body", "error": "Invalid Email"})
		return
	}
	if !utils.IsValidPassword(r.Password) {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse password body", "error": "Invalid Password"})
		return
	}
	var user models.User
	user.Email = strings.ToLower(r.Email)
	fmt.Println("Hashing password ...")
	hashedPassword, err := utils.GetHashedPassword(r.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not hash user password", "error": err.Error()})
		return
	}
	user.Password = hashedPassword
	fmt.Println("New user to save:", user)
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Created user successfully", "user": user})
}

func getUserById(context *gin.Context) {
	fmt.Println("Getting user by id ...")
	type response struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
	}
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id", "error": err.Error()})
		return
	}
	fmt.Println("User id to fetch: ", id)
	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}
	var responseContent response
	responseContent.Email = user.Email
	responseContent.Id = user.ID
	context.JSON(http.StatusOK, gin.H{"message": "user fetched successfully", "user": responseContent})
}

func getUserByEmail(context *gin.Context) {
	fmt.Println("Getting user by email...")
	type response struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	}
	email := context.Param("email")
	fmt.Println("email: ", email)
	user, err := models.GetUserByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user by email", "error": err.Error()})
		return
	}
	var responseContent response
	responseContent.ID = user.ID
	responseContent.Email = user.Email
	context.JSON(http.StatusOK, gin.H{"message": "user fetched successfully", "user": responseContent})
}

func updateUserById(context *gin.Context) {
	fmt.Println("Updating user...")
	type request struct {
		Email string `json:"email" binding:"required"`
	}
	var r request
	var user models.User
	// Parse the user id from the request parameters
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	fmt.Println("  user with id: ", id)
	// Parse the token from the request
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
	if !(id == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	user.ID = id
	err = context.ShouldBindJSON(&r)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "error": err.Error()})
		return
	}
	user.Email = r.Email
	err = user.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated user successfully"})
}

func deleteUserById(context *gin.Context) {
	fmt.Println("Deleting user ...")
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	fmt.Println("User id: ", id)
	// Parse the token from the request
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
	if !(id == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	var user *models.User
	user, err = models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user", "error": err.Error()})
		return
	}
	fmt.Println("Deleting user ...")
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted user successfully"})
}

func changeUserPassword(context *gin.Context) {
	fmt.Println("Changing user password ...")
	type NewPassword struct {
		Value string `json:"value" binding:"required"`
	}
	var newPassword NewPassword
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id", "error": err.Error()})
		return
	}
	fmt.Println("user id:", id)
	// Parse the token from the request
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
	if !(id == userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// fmt.Println("User id to change password: ", id)
	fmt.Println("Parsing request body ...")
	err = context.ShouldBindJSON(&newPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	// fmt.Println("New password: ", newPassword)
	fmt.Println("Changing password ...")
	err = models.ChangeUserPassword(id, newPassword.Value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not change password of user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Changed password successfully"})
}

// Log in the user
func loginUser(context *gin.Context) {
	fmt.Println("Logging in user ...")
	type loginDetails struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var err error
	var login loginDetails
	fmt.Println("Parsing request body ...")
	err = context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse login request data", "error": err.Error()})
		return
	}
	fmt.Println("Validating credentials ...")
	err = models.ValidateCredentials(login.Email, login.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	user, err := models.GetUserByEmail(login.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	fmt.Println("Generating token for user ...")
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user",
			"error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
