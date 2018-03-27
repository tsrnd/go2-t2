package handler

import (
	"go2-t2/config"
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
