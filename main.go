package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Print("HELLO GOLANG")
	log.Println("Server started on: http://localhost:8080")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/detail", detail)
	http.ListenAndServe(":8080", nil)
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
