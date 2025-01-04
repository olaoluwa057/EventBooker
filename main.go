package main

import (
	"example.com/event-booker/db"
	"example.com/event-booker/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}

