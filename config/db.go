package config

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
	"log"
	"os"

	"github.com/jinzhu/gorm"
<<<<<<< HEAD
	_ "github.com/jinzhu/gorm/dialects/postgres"
=======
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
)

//ConnectDB config connect db
func ConnectDB() *gorm.DB {
<<<<<<< HEAD
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
=======
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
	dbConnect := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
<<<<<<< HEAD
	sslMode := os.Getenv("SSLMODE")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, sslMode)
	db, err := gorm.Open(dbConnect, dbInfo)

=======
	db, err := gorm.Open(dbConnect, dbUser+":"+dbPass+"@/"+dbName)
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
	if err != nil {
		log.Fatal(err)
	}
	return db
}
