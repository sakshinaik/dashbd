package main

import (
	"dashboard-api/config"
	"dashboard-api/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) InitApp() {
	db, err := models.MySQLConnect(config.Conf.MySQL_connection)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
	a.Router = mux.NewRouter().StrictSlash(true)
	routes := a.appRoutes()

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		a.Router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)

	}
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))

}
