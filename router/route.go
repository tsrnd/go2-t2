package router

import (
	"go2-t2/handle/blog"
	"net/http"

	"github.com/gorilla/mux"
)

//Route route
func Route() *mux.Router {
	var bh blog.BlogHandler
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/create", bh.Create).Methods("GET")
	r.HandleFunc("/store", bh.Store).Methods("POST")

	return r
}
