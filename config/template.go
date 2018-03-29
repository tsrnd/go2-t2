package config

import "html/template"

// GetTemplate show layout of blog
func GetTemplate(tplName string) *template.Template {
	tplName = "views/blogs/" + tplName
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", tplName, "views/layout/footer.html"))
	return tmpl
}
