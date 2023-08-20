package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	userRequest := struct {
		User struct {
			Username string
			Email    string
			Password string
		}
	}{}

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	as := NewAuthService(db)
	_, err = as.Register(userRequest.User.Username, userRequest.User.Email, userRequest.User.Password)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	u, err := as.Login(userRequest.User.Email, userRequest.User.Password)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			User struct {
				Email    string `json:"email"`
				Token    string `json:"token"`
				Username string `json:"username"`
				Bio      string `json:"bio"`
				Image    string `json:"image"`
			} `json:"user"`
		}{
			User: struct {
				Email    string `json:"email"`
				Token    string `json:"token"`
				Username string `json:"username"`
				Bio      string `json:"bio"`
				Image    string `json:"image"`
			}{
				Email:    u.Email,
				Token:    u.Token.String,
				Username: u.Username,
				Bio:      u.Bio.String,
				Image:    u.Image.String,
			},
		},
	)
}
