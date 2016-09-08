package controller

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/cream-lab/vanilla/app"
	"github.com/cream-lab/vanilla/datastore"
	"github.com/cream-lab/vanilla/model"
)

func AddUser(w http.ResponseWriter, req *http.Request) {
	newUser := model.User{
		Email:        req.FormValue("email"),
		Name:         req.FormValue("name"),
		Password:     req.FormValue("password"),
		IdProvider:   model.GetIdentityProvider(req.FormValue("idp")),
		RegisterDate: time.Now(),
		UpdateDate:   time.Now(),
	}

	datastore, error := datastore.CreateDataStore(app.Config)
	if error != nil {
		log.Fatal(error)
	}

	seq, affectedCount := datastore.AddUser(newUser)
	log.Printf("New user seq: %d, affectedCount: %d", seq, affectedCount)

	var tmpl = template.Must(template.ParseGlob("views/sign_up_completed.html"))
	tmpl.ExecuteTemplate(w, "sign_up_completed.html", nil)
}
