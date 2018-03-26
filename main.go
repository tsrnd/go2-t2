package main

import (
	"fmt"
	"go2-t2/router"
	"log"
	"net/http"
)

func main() {
	fmt.Print("HELLO GOLANG")
	log.Println("Server started on: http://localhost:8080")

	r := router.Route()

	http.ListenAndServe(":8080", r)
}
