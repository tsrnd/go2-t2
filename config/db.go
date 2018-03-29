package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

//ConnectDB config connect db
func ConnectDB() *gorm.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	db, err := gorm.Open(dbConnect, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
