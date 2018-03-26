package controller

import (
	"go2-t2/model"
	"html/template"
	"net/http"
)

// BlogController blog controller
type BlogController struct {
}

//Create create blog layout
func (bc BlogController) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/create.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "create", nil)
}

//SaveBlog save blog
func (bc BlogController) SaveBlog(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	model.CreateBlog(title, content)
}
