package main

import "testing"

func TestRegister(t *testing.T) {
	db := NewSeededInMemDB()
	as := NewAuthService(db)

	username := "neueruser"
	email := "neu@user.de"
	password := "neuneu123!"

	u, err := as.Register(username, email, password)
	if err != nil {
		t.Fatalf("register err %#v", err)
	}
	if u.ID == 0 {
		t.Errorf("user id %#v", u.ID)
	}

	if u.Username != username {
		t.Errorf("username should %#v but is %#v", username, u.Username)
	}

	if u.Email != email {
		t.Errorf("email should %#v but is %#v", email, u.Email)
	}

	if u.Password == "" || u.Password == password {
		t.Errorf("user password is empty or not hashed %#v", u.Password)
	}
}

func TestLogin(t *testing.T) {
	db := NewSeededInMemDB()
	as := NewAuthService(db)

	username := "neueruser"
	email := "neu@user.de"
	password := "neuneu123!"

	_, err := as.Register(username, email, password)
	if err != nil {
		t.Fatalf("register err %#v", err)
	}

	u, err := as.Login(email, password)
	if err != nil {
		t.Fatalf("register err %#v", err)
	}

	if u.Username != username {
		t.Errorf("username should %#v but is %#v", username, u.Username)
	}

	if u.Email != email {
		t.Errorf("email should %#v but is %#v", email, u.Email)
	}

	if u.Password == "" || u.Password == password {
		t.Errorf("user password is empty or not hashed %#v", u.Password)
	}

	if u.Token.String == "" || !u.Token.Valid {
		t.Errorf("user Token is empty %#v", u.Token)
	}

}
