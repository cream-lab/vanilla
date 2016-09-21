package controller

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, req *http.Request) {
	var tmpl = template.Must(template.ParseGlob("views/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
