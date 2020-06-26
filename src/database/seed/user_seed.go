package seed

import (
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is seeder to create user
func CreateUser(db *gorm.DB) error {
	user := entity.User{}
	if err := faker.FakeData(&user); err != nil {
		logrus.Errorln("Error user seed", err)
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	hashedPassword := string(bytes)

	return db.Create(&entity.User{
		Name:     user.Name,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
		Password: hashedPassword,
	}).Error
}
