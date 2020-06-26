package services

import (
	"github.com/arifseft/go-auth/src/auth"
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/arifseft/go-auth/src/middlewares/exception"
	"github.com/arifseft/go-auth/src/repositories"
	"github.com/arifseft/go-auth/src/validations/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository  repositories.UserRepository
	LoginRepository repositories.LoginRepository
}

func LService() LoginService {
	return LoginService{
		LoginRepository: repositories.LRepository(),
		UserRepository:  repositories.URepository(),
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type LoginResponse struct {
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func (s *LoginService) Login(login schemas.Login) *LoginResponse {
	userExist := s.UserRepository.UserExist(
		repositories.UserExistParams{Email: login.Username},
	)

	if (userExist == entity.User{}) {
		exception.NotFound("User not exist", []map[string]interface{}{
			{"message": "User with this email / username not found", "flag": "USER_NOT_FOUND"},
		})
	}

	password, _ := hashPassword(login.Password)
	match := checkPasswordHash(userExist.Email, password)

	if match {
		exception.BadRequest("Unauthorized", []map[string]interface{}{
			{"message": "Wrong password", "flag": "USER_UNATHORIZED"},
		})
	}

	authData, err := s.LoginRepository.CreateAuth(userExist.ID)
	if err != nil {
		exception.Unauthorized("Unauthorized", []map[string]interface{}{
			{"message": "User unauthorized", "flag": "USER_UNATHORIZED"},
		})
	}

	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID

	token, loginErr := auth.CreateToken(authD)
	if loginErr != nil {
		exception.Forbidden("Please try to login latter", []map[string]interface{}{
			{"message": "Status forbidden", "flag": "USER_FORBIDDEN"},
		})
	}

	data := &LoginResponse{
		UserID:      authD.UserId,
		Email:       login.Username,
		AccessToken: token,
	}

	return data
}

func (s *LoginService) Logout(c *gin.Context) string {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		exception.Forbidden("Please try to login later", []map[string]interface{}{
			{"message": "Status forbidden", "flag": "USER_FORBIDDEN"},
		})
	}
	err = s.LoginRepository.DeleteAuth(au)
	if err != nil {
		exception.Unauthorized("Unauthorized", []map[string]interface{}{
			{"message": "Unauthorized", "flag": "USER_UNATHORIZED"},
		})
	}

	return "success"
}
