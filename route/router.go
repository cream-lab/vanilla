package route

import (
	"github.com/cream-lab/vanilla/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Home)
	
	return router
}
