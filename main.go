package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/cream-lab/vanilla/app"
	"github.com/cream-lab/vanilla/datastore"
	"github.com/cream-lab/vanilla/route"
)

//
// type Env struct {
// 	db *datastore.DataStore
// }

var Conf = make(map[string]interface{})

func main() {
	app.LoadConfig()
	datastore.Register("inmemory", datastore.NewInMemoryDataStore)
	// db, error := datastore.CreateDataStore(conf)

	router := route.Router()

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8080", n)
}
