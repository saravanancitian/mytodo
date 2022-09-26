package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var parsedTemplate *template.Template

func iniweb() {
	parsedTemplate = template.Must(template.ParseFiles("static/index.html", "static/signin.html", "static/signup.html"))
}

func servhtml(w http.ResponseWriter, r *http.Request) {
	htmlfile := r.URL.Path[len("/"):]

	if len(htmlfile) > 0 {
		parsedTemplate.ExecuteTemplate(w, htmlfile, nil)
	} else {
		parsedTemplate.ExecuteTemplate(w, "index.html", nil)
	}
}

func servJs(w http.ResponseWriter, r *http.Request) {
	jsfile := r.URL.Path[len("/js/"):]

	http.ServeFile(w, r, "static/js/"+jsfile)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var usrid int
	body, err := io.ReadAll(r.Body)
	usr := struct {
		Name        string `json:"name,omitempty"`
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
		ConfirmPass string `json:"confirmPass,omitempty"`
	}{}

	if err != nil {
		w.Write([]byte("failed"))
		return
	}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		w.Write([]byte("failed"))
		return
	}

	//if usr.Password == usr.ConfirmPass {
	row := db.QueryRow("select userid from user where email=?", usr.Email)
	err = row.Scan(&usrid)
	if err == sql.ErrNoRows {
		_, err = insert_user.Exec(usr.Name, usr.Email, usr.Password)
		if err != nil {

			w.Write([]byte("Error Registering"))
		} else {
			w.Write([]byte("Success"))
		}
	} else {
		w.Write([]byte("user already exist"))
	}

	// } else {
	// 	w.Write([]byte("Confirm password not match"))
	// }

}

func login(w http.ResponseWriter, r *http.Request) {
	var qpass string

	body, err := io.ReadAll(r.Body)
	usr := user{}

	if err != nil {
		w.Write([]byte("failed"))
		return
	}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		w.Write([]byte("failed"))
		return
	}

	rows := db.QueryRow("select password from user where email=?", usr.Email)

	rows.Scan(&qpass)
	if usr.Password == qpass {
		w.Write([]byte("Success"))

	} else {
		w.Write([]byte("Failure"))
	}
}
