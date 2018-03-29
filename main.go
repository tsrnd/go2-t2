package main

import (
	"go2-t2/config"
	"go2-t2/model"
	"go2-t2/router"
	"log"
	"net/http"
	"os"
<<<<<<< HEAD

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
=======
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
)

func main() {
	log.Printf("Server started on: http://localhost%s", os.Getenv("SERVER_PORT"))

	r := router.Route()
	http.ListenAndServe(os.Getenv("SERVER_PORT"), context.ClearHandler(r))
}

func init() {
	config.SetEnv()
	db := config.ConnectDB()
	model.SetDatabase(db)
}
