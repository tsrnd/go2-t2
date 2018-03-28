package handler

import (
	"go2-t2/config"
	"go2-t2/model"
	"net/http"
)

type BlogHandler struct{}

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("index.html")
	tmpl.ExecuteTemplate(w, "index", nil)
}

func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("edit.html")
	tmpl.ExecuteTemplate(w, "edit", nil)
}

func (bh BlogHandler) Detail(w http.ResponseWriter, r *http.Request) {
	tmpl := config.GetTemplate("detail.html")
	tmpl.ExecuteTemplate(w, "detail", nil)
}

func (bh BlogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")
	db := model.DBCon
	db.Delete(model.Blog{}, "id = ?", postID)

	http.Redirect(w, r, "/", 301)
}
