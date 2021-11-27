package app

import (
	"github.com/gorilla/mux"
	"url-shortener/app/database"
)

type App struct {
	Router *mux.Router
	DB     database.UrlShortDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/shorten", a.CreateShortUrlHandler()).Methods("POST")
	a.Router.HandleFunc("/{shortUrl}", a.GetShortUrlHandler()).Methods("GET")
}
