package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	// middlewares must come before routes

	// routes must be here
	app.App.Routes.Get("/", app.Handlers.Home)

	// static files must be there
	fileServer := http.FileServer(http.Dir("./public"))
	app.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return app.App.Routes
}
