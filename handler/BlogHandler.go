package handler

import (
	"go2-t2/config"
	"go2-t2/model"
	"go2-t2/repository"
	"net/http"

	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

type BlogHandler struct{}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//Index show list blogs
func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	blogs := repository.GetAllPosts()
	tmpl := config.GetTemplate("index.html")
	tmpl.ExecuteTemplate(w, "index", blogs)
}

//Edit form edit blog layout
func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog := repository.Get(id)

	resource := make(map[string]interface{})
	resource["blog"] = blog

	tmpl := config.GetTemplate("edit.html")
	tmpl.ExecuteTemplate(w, "edit", resource)
}

//Update update blog
func (bh BlogHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog := model.Blog{Title: r.FormValue("title"), Content: r.FormValue("content")}
	err := validate.Struct(blog)
	if err != nil {
		errors := make(map[string]interface{})

		for _, errV := range err.(validator.ValidationErrors) {
			errors[errV.Field()] = errV.Field() + " " + errV.ActualTag()
		}

		blog := repository.Get(id)

		resource := make(map[string]interface{})
		resource["blog"] = blog
		resource["error"] = errors

		tmpl := config.GetTemplate("edit.html")
		tmpl.ExecuteTemplate(w, "edit", resource)
		return
	}

	model.DBCon.Table("blog_golang.blogs").Where("id = (?)", id).Updates(blog)
	http.Redirect(w, r, "/", 301)
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
		errors := make(map[string]interface{})

		for _, errV := range err.(validator.ValidationErrors) {
			errors[errV.Field()] = errV.Field() + " " + errV.ActualTag()
		}

		tmpl := config.GetTemplate("create.html")
		tmpl.ExecuteTemplate(w, "create", errors)
	} else {
		repository.Store(blog)

		http.Redirect(w, r, "/", 301)
	}
}
