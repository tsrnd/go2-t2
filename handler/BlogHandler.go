package handler

import (
	"go2-t2/config"
	"go2-t2/model"
	"net/http"

	"github.com/go-chi/chi"
)

type BlogHandler struct{}

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("index.html")
	tmpl.ExecuteTemplate(w, "index", nil)
}

//Edit form edit blog layout
func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog := model.Blog{}
	db := model.DBCon
	db.First(&blog, id)
	tmpl := config.GetTemplate("edit.html")
	tmpl.ExecuteTemplate(w, "edit", blog)
}

//Update update blog
func (bh BlogHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog := model.Blog{Title: r.FormValue("title"), Content: r.FormValue("content")}
	model.DBCon.Table("blogs").Where("id IN (?)", id).Updates(blog)
	model.Redirect(w, r)
}

func (bh BlogHandler) Detail(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("detail.html")
	tmpl.ExecuteTemplate(w, "detail", nil)
}

//Create create blog layout
func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("create.html")
	tmpl.ExecuteTemplate(w, "create", nil)
}

//Store save blog
func (bh BlogHandler) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	// model.CreateBlog(title, content)
	// model.Redirect(w, r)
	blog := model.Blog{Title: title, Content: content}
	db := model.DBCon
	db.NewRecord(blog) // => returns `true` as primary key is blank
	db.Create(&blog)
	model.Redirect(w, r)
}
