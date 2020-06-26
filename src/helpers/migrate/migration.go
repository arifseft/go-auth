package main

import (
	db "github.com/arifseft/go-auth/src/database"
	"github.com/arifseft/go-auth/src/database/migration"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.AppConnection()
	conn := db.GetDB()
	defer conn.Close()

	migration.CreateUser(conn)
}
