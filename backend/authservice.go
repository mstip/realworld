package main

import (
	"database/sql"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Email    string
	Username string
	Bio      sql.NullString
	Image    sql.NullString
	Token    sql.NullString
	Password string
}

type AuthService struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) Register(username string, email string, password string) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, err
	}

	_, err = as.db.Exec(`INSERT INTO users(username, email, password) values(?,?,?)`, username, email, hashedPassword)
	if err != nil {
		return User{}, err
	}

	var u []User

	err = as.db.Select(&u, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return User{}, err
	}

	if len(u) == 0 {
		return User{}, errors.New("failed to create user, cant select new user")
	}

	return u[0], nil
}

func (as *AuthService) Login(email string, password string) (User, error) {
	var dbUsers []User

	err := as.db.Select(&dbUsers, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return User{}, err
	}

	if len(dbUsers) == 0 {
		return User{}, errors.New("failed to find user with email: " + email)
	}

	u := dbUsers[0]

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	jwt, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte("SUPER SECRET"))
	if err != nil {
		return User{}, err
	}

	u.Token.String = jwt
	u.Token.Valid = true

	_, err = as.db.Exec(`UPDATE users SET token = ? WHERE id = ?`, jwt, u.ID)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
