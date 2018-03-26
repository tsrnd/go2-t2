package main

import (
	"fmt"
	"go2-t2/config"
	"go2-t2/model"
	"go2-t2/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Print("HELLO GOLANG")

	log.Printf("Server started on: http://localhost%s", os.Getenv("SERVER_PORT"))

	r := router.Route()

	http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}

func init() {
	config.SetEnv()
	db := config.ConnectDB()
	model.SetDatabase(db)
}
