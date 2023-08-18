package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

func AllTagsHanlder(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	tagService, _ := NewTagService(db)
	tags, err := tagService.AllTags()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			Tags []string `json:"tags"`
		}{
			Tags: tags,
		},
	)
}

func NewRouter(db *sqlx.DB) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) { AllTagsHanlder(w, r, db) })

	return cors.Default().Handler(r)
}
