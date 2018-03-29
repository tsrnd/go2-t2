package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ConnectDB config connect db
func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbConnect := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	sslMode := os.Getenv("SSLMODE")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, sslMode)
	db, err := gorm.Open(dbConnect, dbInfo)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
