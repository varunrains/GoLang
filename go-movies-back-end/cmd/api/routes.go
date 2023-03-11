package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//create a router multiplexer

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Get("/", app.Home)

	return mux
}
