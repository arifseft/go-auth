package controllers

import (
	"net/http"

	"github.com/arifseft/go-auth/src/services"
	"github.com/arifseft/go-auth/src/utils/flag"
	"github.com/arifseft/go-auth/src/validations/schemas"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginController struct {
	Service services.LoginService
}

func LController() LoginController {
	return LoginController{
		Service: services.LService(),
	}
}

// Login -> login route
// POST /login
func (l *LoginController) Login(c *gin.Context) {
	var login schemas.Login
	_ = c.ShouldBindBodyWith(&login, binding.JSON)

	data := l.Service.Login(login)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.LoginSuccess.Message,
		"data":    data,
		"error":   nil,
	})
}

func (l *LoginController) Logout(c *gin.Context) {
	data := l.Service.Logout(c)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.LogoutSuccess.Message,
		"data":    data,
		"error":   nil,
	})
}
