package handler

import (
	"go2-t2/config"
	"go2-t2/handler/repository"
	"go2-t2/model"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type BlogHandler struct{}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("index.html")
	tmpl.ExecuteTemplate(w, "index", nil)
}

func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("edit.html")
	tmpl.ExecuteTemplate(w, "edit", nil)
}

//Create create blog layout
func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("create.html")
	tmpl.ExecuteTemplate(w, "create", nil)
}

//Store save blog
func (bh BlogHandler) Store(w http.ResponseWriter, r *http.Request) {
	blog := model.Blog{Title: r.FormValue("title"), Content: r.FormValue("content")}
	err := validate.Struct(blog)
	if err != nil {
		tmpl := config.GetTemplate("create.html")
		tmpl.ExecuteTemplate(w, "create", err)
	} else {
		repository.Store(blog)

		http.Redirect(w, r, "/", 301)
	}
}
