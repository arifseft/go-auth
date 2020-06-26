package auth

import (
	"net/http"

	"github.com/arifseft/go-auth/src/auth"
	"github.com/arifseft/go-auth/src/utils/flag"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": flag.LogoutUnauthorized.Message,
				"data":    nil,
				"error":   nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
