package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booker/modals"
	"example.com/event-booker/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := modals.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events"})
	}
	context.JSON(http.StatusOK, gin.H{"Events": events})
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	_, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var event modals.Event

	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	event.User_ID = context.GetInt64("user_id")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func getEvent(context *gin.Context) {
	id := context.Param("id")

	eventID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	event, err := modals.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")

	eventID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	event, err := modals.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "id not found"})
		return
	}

	userId := context.GetInt64("user_id")

	users, err := modals.GetUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get user"})
	}

	if userId != event.User_ID && !users.IsAdmin {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to update event"})
		return
	}

	err = event.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")

	eventID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	event, err := modals.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}

	userId := context.GetInt64("user_id")

	if userId != event.User_ID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func registerForEvent(context *gin.Context) {

	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	event, err := modals.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}

	userId := context.GetInt64("user_id")

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		return
	}

	event, err := modals.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event"})
		return
	}

	userId := context.GetInt64("user_id")
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully"})
}
