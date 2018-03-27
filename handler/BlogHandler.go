package handler

import (
	"go2-t2/model"
	"html/template"
	"net/http"
)

type BlogHandler struct{}

func getTemplate(str string) *template.Template {
	str = "views/blogs/" + str
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", str, "views/layout/footer.html"))
	return tmpl
}

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	blogs := model.GetAllPosts()
	tmpl := getTemplate("index.html")
	tmpl.ExecuteTemplate(w, "index", blogs)
}

func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	tmpl := getTemplate("edit.html")
	tmpl.ExecuteTemplate(w, "edit", nil)
}

func (bh BlogHandler) Detail(w http.ResponseWriter, r *http.Request) {
	tmpl := getTemplate("detail.html")
	tmpl.ExecuteTemplate(w, "detail", nil)
}
