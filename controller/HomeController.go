package HomeController

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))

func Home(w http.ResponseWriter, req *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
