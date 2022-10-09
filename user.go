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
	var usr User
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		error := Error{1, err.Error()}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	rows := db.QueryRow("select * from user where email=?", email)

	err = rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)
	if err != nil {
		error := Error{1, err.Error()}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
	}
	if password == usr.Password {
		gotoProjectPg(&usr, w, r)

	} else {
		error := Error{1, "Incorrect user id or password"}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	var usrid int

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		error := Error{1, err.Error()}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPass := r.FormValue("confirmPass")

	if password == confirmPass {
		row := db.QueryRow("select userid from user where email=?", email)

		err = row.Scan(&usrid)
		if err == sql.ErrNoRows {
			_, err = insert_user.Exec(name, email, password)
			if err != nil {

				error := Error{1, err.Error()}
				parsedTemplate.ExecuteTemplate(w, "error.html", error)

			} else {
				rows := db.QueryRow("select * from user where email=?", email)
				var usr User

				err = rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)
				if err != nil {
					error := Error{1, err.Error()}
					parsedTemplate.ExecuteTemplate(w, "error.html", error)
				}
				gotoProjectPg(&usr, w, r)
			}
		} else {
			error := Error{1, "user already exist"}
			parsedTemplate.ExecuteTemplate(w, "error.html", error)
		}

	} else {
		error := Error{1, "Confirm password not match"}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
	}
}
