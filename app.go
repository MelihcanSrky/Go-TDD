package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type App struct {
	Router *chi.Mux
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbName string) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)

	var err error

	a.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = chi.NewRouter()
}

func (a *App) Run(addr string) {}
