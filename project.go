package main

import (
	"log"
	"net/http"
	"strconv"
)

type Project struct {
	ProjectId int
	Name      string
	UserId    int
	Completed bool
	des       string
}

func getProjectList(userid int) []Project {
	var projects []Project
	rows, err := db.Query("Select * from project where userid = ?", userid)
	if err != nil {
		//TODO need check for error
		log.Fatal("Error in query")
	}
	defer rows.Close()

	for rows.Next() {
		var project Project
		err = rows.Scan(&project.ProjectId, &project.Name, &project.UserId, &project.Completed, &project.des)
		if err != nil {

			break
		}
		projects = append(projects, project)
	}

	return projects

}

func gotoProjectPg(usr *User, w http.ResponseWriter, r *http.Request) {

	var projects []Project = getProjectList(usr.Userid)

	if len(projects) > 0 {
		var data = struct {
			Username string
			Projects []Project
		}{usr.Name, projects}
		parsedTemplate.ExecuteTemplate(w, "project_list.html", data)

	} else {
		var data = struct {
			Username string
			Userid   int
		}{usr.Name, usr.Userid}
		err := parsedTemplate.ExecuteTemplate(w, "create_project.html", data)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func create_project(w http.ResponseWriter, r *http.Request) {
	var userid int
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form", err)
		error := Error{1, err.Error()}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)
		return
	}

	name := r.FormValue("name")
	desc := r.FormValue("desc")
	userid, _ = strconv.Atoi(r.FormValue("userid"))
	username := r.FormValue("username")
	_, err = insert_project.Exec(name, userid, false, desc)
	if err != nil {

		error := Error{1, err.Error()}
		parsedTemplate.ExecuteTemplate(w, "error.html", error)

	} else {
		var projects []Project = getProjectList(userid)

		var data = struct {
			Username string
			Projects []Project
		}{username, projects}
		parsedTemplate.ExecuteTemplate(w, "project_list.html", data)

	}

}
