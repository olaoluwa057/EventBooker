package middleware

import (
	"net/http"
	"example.com/event-booker/utils"
	"github.com/gin-gonic/gin"
)

func Authorization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user_id, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("user_id", user_id)
	context.Next()
}
