package routes

import (
	"net/http"

	"github.com/arifseft/go-auth/src/controllers"
	"github.com/arifseft/go-auth/src/middlewares/auth"
	"github.com/arifseft/go-auth/src/validations"
	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {

	loginController := controllers.LController()
	{
		g.POST("/login", loginController.Login)
		g.POST("/logout", auth.TokenAuthMiddleware(), loginController.Logout)
	}

	userController := controllers.UController()
	{
		g.GET("/users", userController.GetUsers)
		g.GET("/user/:id", validations.GetUser, userController.GetUser)
		g.POST("/user", validations.CreateUser, userController.CreateUser)
		g.PATCH("/user/:id", validations.UpdateUser, userController.UpdateUser)
		g.DELETE("/user/:id", validations.DeleteUser, userController.DeleteUser)

		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "ready",
				"database": "error",
			})
		})
	}
}
