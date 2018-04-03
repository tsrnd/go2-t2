package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	mMiddleware "github.com/tsrnd/trainning/shared/middleware"
	"github.com/tsrnd/trainning/user"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {

	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// user set.
	uh := user.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)
	r.Mux.Route("/v1", func(cr chi.Router) {
		cr.Post("/register/device", uh.RegisterByDevice)
	})
}
