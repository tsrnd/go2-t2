package router

import (
	"go2-t2/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func Route() *chi.Mux {
	var bh handler.BlogHandler
	r := chi.NewRouter()

	r.Method(http.MethodGet, "/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Get("/", bh.Index)
	r.Get("/{id}/edit", bh.Edit)
	r.Get("/detail", bh.Detail)
	r.Post("/delete/{{.ID}}", bh.Delete)
	return r
}
