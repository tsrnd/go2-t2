package main

import (
	"go2-t2/config"
	"go2-t2/model"
	"go2-t2/router"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Server started on: http://localhost%s", os.Getenv("SERVER_PORT"))

	r := router.Route()
	http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}

func init() {
	config.SetEnv()
	db := config.ConnectDB()
	model.SetDatabase(db)
}
