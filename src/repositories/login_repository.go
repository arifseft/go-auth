package repositories

import (
	"github.com/arifseft/go-auth/src/auth"
	db "github.com/arifseft/go-auth/src/database"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

type LoginRepository struct {
	Conn *gorm.DB
}

func LRepository() LoginRepository {
	return LoginRepository{Conn: db.GetDB().Table("auths")}
}

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint   `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

func (r *LoginRepository) CreateAuth(userId uint) (*Auth, error) {
	au := &Auth{}
	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.UserID = userId
	err := r.Conn.Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (r *LoginRepository) DeleteAuth(authD *auth.AuthDetails) error {
	au := &Auth{}
	db := r.Conn.Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
