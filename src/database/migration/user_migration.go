package migration

import (
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CreateUser is create user tabel for migration
func CreateUser(conn *gorm.DB) {
	conn.AutoMigrate(&entity.User{})

	logrus.Info("Success running migration")
}
