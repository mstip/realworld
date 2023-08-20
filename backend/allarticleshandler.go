package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func AllArticlesHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	articleService := NewArticleService(db)
	articles, err := articleService.AllArticles()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			Articles []Article `json:"articles"`
		}{
			Articles: articles,
		},
	)
}
