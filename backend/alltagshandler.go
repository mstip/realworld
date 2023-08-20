package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func AllTagsHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
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
