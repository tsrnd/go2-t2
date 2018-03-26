package router

import (
	"go2-t2/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	var bc controller.BlogController
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/", bc.Index).Methods("GET")
	r.HandleFunc("/edit", bc.Edit).Methods("GET")
	r.HandleFunc("/detail", bc.Detail).Methods("GET")

	return r
}
