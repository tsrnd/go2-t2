package config

import "html/template"

func GetTemplate(str string) *template.Template {
	str = "views/blogs/" + str
	tmpl := template.Must(template.ParseFiles("views/layout/header.html", str, "views/layout/footer.html"))
	return tmpl
}
