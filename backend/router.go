package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

func NewRouter(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) { AllTagsHandler(w, r, db) }).Methods(http.MethodGet)
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) { RegisterHandler(w, r, db) }).Methods(http.MethodPost)
	r.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) { LoginHandler(w, r, db) }).Methods(http.MethodPost)
	r.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) { AllArticlesHandler(w, r, db) }).Methods(http.MethodGet)
	return cors.Default().Handler(r)
}
