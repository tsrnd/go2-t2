package blog

import (
	"go2-t2/model"
	"html/template"
	"net/http"
)

//BlogHandler blog struct
type BlogHandler struct {
}

//Create create blog layout
func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", "views/blogs/create.html", "views/layout/footer.html"))
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
