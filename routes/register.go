package routes

import (
	"example.com/event-booker/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/users", createUser)
	server.POST("login", login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authorization)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.POST("/events/:id/cancelregistration", cancelRegistration)

}
