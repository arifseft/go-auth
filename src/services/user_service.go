package services

import (
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/arifseft/go-auth/src/middlewares/exception"
	"github.com/arifseft/go-auth/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserService -> the propose of user service is handling business logic application
type UserService struct {
	UserRepository repositories.UserRepository
}

// UService -> user service instance
func UService() UserService {
	return UserService{
		UserRepository: repositories.URepository(),
	}
}

// GetUsers -> get users service logic
func (s *UserService) GetUsers() []repositories.GetUser {
	users := s.UserRepository.GetUsers()
	return users
}

// GetUser -> get user service logic
func (s *UserService) GetUser(id int64) repositories.GetUser {
	user := s.UserRepository.GetUser(id)

	if (user == repositories.GetUser{}) {
		exception.NotFound("User not found", []map[string]interface{}{
			{"message": "User with this ID not found", "flag": "USER_NOT_FOUND"},
		})
	}

	return user
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser -> create user service logic
func (s *UserService) CreateUser(user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(
		repositories.UserExistParams{Email: user.Email},
	)

	if (userExist != entity.User{}) {
		exception.Conflict("User conflict", []map[string]interface{}{
			{"message": "User with this email already exist", "flag": "USER_ALREADY_EXIST"},
		})
	}

	user.Password, _ = hashPassword(user.Password)
	data := s.UserRepository.CreateUser(user)
	return data
}

// UpdateUser -> update user service logic
func (s *UserService) UpdateUser(id uint, user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.NotFound("User not exist", []map[string]interface{}{
			{"message": "User with this ID not found", "flag": "USER_NOT_FOUND"},
		})
	}

	data := s.UserRepository.UpdateUser(id, user)
	return data
}

// DeleteUser -> delete user service logic
func (s *UserService) DeleteUser(id uint) repositories.GetUser {
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.NotFound("User not exist", []map[string]interface{}{
			{"message": "User with this ID not found", "flag": "USER_NOT_FOUND"},
		})
	}

	data := s.UserRepository.DeleteUser(id)
	return data
}
