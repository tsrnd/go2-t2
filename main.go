package main

import (
	"fmt"
	"go2-t2/config"
	"go2-t2/model"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Print("HELLO GOLANG")
	log.Printf("Server started on: http://localhost%s", os.Getenv("SERVER_PORT"))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/detail", detail)
	http.ListenAndServe(os.Getenv("SERVER_PORT"), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/index.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "index", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/edit.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "edit", nil)
}

func detail(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/detail.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "detail", nil)
}

func init() {
	config.SetEnv()
	db := config.ConnectDB()
	model.SetDatabase(db)
}
