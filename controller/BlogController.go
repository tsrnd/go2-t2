package controller

import (
	"go2-t2/model"
	"html/template"
	"net/http"
)

type (
	BlogController struct{}
)

func (bc BlogController) Index(w http.ResponseWriter, r *http.Request) {
	blogs := model.GetAllPosts()
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/index.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "index", blogs)
}

func (bc BlogController) Edit(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/edit.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "edit", nil)
}

func (bc BlogController) Detail(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/detail.html", "views/layout/footer.html"))
	tmpl.ExecuteTemplate(w, "detail", nil)
}
