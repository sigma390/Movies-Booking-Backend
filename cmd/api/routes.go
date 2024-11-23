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
	mux.Get("/", app.Home);
	return mux;
}