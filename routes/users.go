package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booker/modals"
	"example.com/event-booker/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user modals.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func login(context *gin.Context) {

	var user modals.User
	var token string

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err, userValue := user.Validate()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	if userValue.IsAdmin {
		token, err = utils.GenerateAdminToken(user.EMAIL, user.ID)

	} else {
		token, err = utils.GenerateToken(user.EMAIL, user.ID)

	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "code": http.StatusOK, "token": token})

}

func getUser(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	user, err := modals.GetUser(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get user"})
		return
	}

	userId := context.GetInt64("user_id")

	users, err := modals.GetUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get user"})
	}

	if userId != users.ID || users.IsAdmin == false {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to get event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}
