package main

import (
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Error struct {
	ErrorNo  int
	ErrorTxt string
}

var parsedTemplate *template.Template

var servermux *http.ServeMux

func iniweb() {
	parsedTemplate = template.Must(template.ParseFiles("static/index.html", "static/signin.html", "static/signup.html", "static/error.html",
		"static/create_project.html", "static/project_list.html", "static/add_update_task.html", "static/task_list.html"))

	servermux = http.NewServeMux()
	servermux.HandleFunc("/js/", servJs)

	servermux.HandleFunc("/", servhtml)
	servermux.HandleFunc("/*.html", servhtml)

	servermux.HandleFunc("/login", login)
	servermux.HandleFunc("/signup", signup)
	servermux.HandleFunc("/project/create", create_project)
	servermux.HandleFunc("/project/", getProject)
	servermux.HandleFunc("/task/", getTask)

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
