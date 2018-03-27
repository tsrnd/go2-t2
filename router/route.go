package router

import (
	"go2-t2/handler"
	"net/http"

	"github.com/go-chi/chi"
)

//Route route
func Route() *chi.Mux {
	var bh handler.BlogHandler
	r := chi.NewRouter()

	r.Method(http.MethodGet, "/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Get("/create", bh.Create)
	r.Post("/store", bh.Store)
	r.Get("/", bh.Index)
	r.Get("/{id}/edit", bh.Edit)

	return r
}
