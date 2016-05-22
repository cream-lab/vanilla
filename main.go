package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/cream-lab/vanilla/route"
)

func main() {
	router := route.Router()

	n := negroni.Classic()
  n.Use(negroni.NewStatic(http.Dir("bower_components/jquery/dist")))
  n.Use(negroni.NewStatic(http.Dir("bower_components/bootstrap/dist")))
	n.UseHandler(router)

	http.ListenAndServe(":8080", n)
}
