package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	Userid   int    `json:"userid,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var usr *User
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		error := Error{1, err.Error()}
		showError(error, w)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	usr, err = fetchUserForEmail(email)
	if err != nil {
		error := Error{1, err.Error()}
		showError(error, w)
	}

	if password == usr.Password {
		gotoProjectPg(usr, w, r)

	} else {
		error := Error{1, "Incorrect user id or password"}
		showError(error, w)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		error := Error{1, err.Error()}
		showError(error, w)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPass := r.FormValue("confirmPass")

	if password == confirmPass {
		var usr *User

		usr, err = fetchUserForEmail(email)

		if err == sql.ErrNoRows {
			usr = new(User)
			usr.Name = name
			usr.Email = email
			usr.Password = password
			err = saveUser(usr)

			if err != nil {

				error := Error{1, err.Error()}
				showError(error, w)

			} else {

				usr, err = fetchUserForEmail(email)

				if err != nil {
					error := Error{1, err.Error()}
					showError(error, w)
					return
				}
				gotoProjectPg(usr, w, r)
			}
		} else {
			error := Error{1, "user already exist"}
			showError(error, w)
		}

	} else {
		error := Error{1, "Confirm password not match"}
		showError(error, w)
	}
}
