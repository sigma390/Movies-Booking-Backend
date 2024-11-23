package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

 func (app *application) routes() http.Handler{

	//create A mux
	mux:=chi.NewRouter();
	//add Middlewares here 
	mux.Use(middleware.Recoverer);
	mux.Use(app.enableCors);
	mux.Get("/", app.Home);
	//get All Movies
	mux.Get("/movies", app.AllMovies);
	return mux;
}