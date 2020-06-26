package migration

import (
	"github.com/arifseft/go-auth/src/database/entity"
	"github.com/jinzhu/gorm"
)

// AutoMigration is auto migrate database
func AutoMigration(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{})
}
