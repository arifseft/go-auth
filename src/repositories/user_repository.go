package repositories

import (
	"time"

	db "github.com/arifseft/go-auth/src/database"
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/jinzhu/gorm"
)

// UserRepository -> the propose of user repository is handling query for user entity
type UserRepository struct {
	Conn *gorm.DB
}

// URepository -> user repository instance to get user table connection
func URepository() UserRepository {
	return UserRepository{Conn: db.GetDB().Table("users")}
}

// GetUser -> get user struct format
type GetUser struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Age       int64      `json:"age,omitempty"`
	Address   string     `json:"address,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// GetUsers -> method to get all users in database
func (r *UserRepository) GetUsers() []GetUser {
	var users []GetUser
	r.Conn.Select("id, name, email, address, age").Find(&users)
	return users
}

// GetUser -> method to get specific user by id
func (r *UserRepository) GetUser(id int64) GetUser {
	user := GetUser{}
	r.Conn.Select("id, name, email, address, age").Where("id = ?", id).First(&user)
	return user
}

// UserExistParams -> Optional params for user exist
type UserExistParams struct {
	Email string
	ID    uint
}

// UserExist -> method to check if user already exist in database by email or id
func (r *UserRepository) UserExist(param UserExistParams) entity.User {
	user := entity.User{}
	if param.ID == 0 {
		r.Conn.Select("id, email, password").Where(&entity.User{Email: param.Email}).First(&user)
	} else {
		r.Conn.Select("id, email, password").Where(&entity.User{ID: param.ID}).First(&user)
	}
	return user
}

// CreateUser -> method to add user in database
func (r *UserRepository) CreateUser(user entity.User) GetUser {
	r.Conn.Create(&user)
	userCreated := GetUser{}
	r.Conn.Select("id, name, email, address, age").Where("id = ?", user.ID).First(&userCreated)
	return userCreated
}

// UpdateUser -> method to update user by id
func (r *UserRepository) UpdateUser(id uint, update entity.User) GetUser {
	r.Conn.Where("id = ?", id).Updates(update)
	userUpdated := GetUser{}
	r.Conn.Select("id, name, email, address, age").Where("id = ?", id).First(&userUpdated)
	return userUpdated
}

// DeleteUser -> method to delete user by id
func (r *UserRepository) DeleteUser(id uint) GetUser {
	deletedUser := GetUser{}
	r.Conn.Select("id, name, email, address, age").Where("id = ?", id).First(&deletedUser)
	r.Conn.Where("id = ?", id).Delete(entity.User{})
	return deletedUser
}
