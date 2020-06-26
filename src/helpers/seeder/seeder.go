package main

import (
	db "github.com/arifseft/go-auth/src/database"
	"github.com/arifseft/go-auth/src/database/seed"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	db.SeedConnection()
	conn := db.GetDB()
	defer conn.Close()

	for _, seed := range seed.All() {
		if err := seed.Run(conn); err != nil {
			logrus.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
