package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	userID, err := ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized%v", "err": err.Error()})
		c.Abort()
		return
	}
	//c.Set("role", role)
	c.Set("userId", userID)
	c.Next()
}
