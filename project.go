package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Project struct {
	ProjectId int
	Name      string
	UserId    int
	Completed bool
	des       string
}

func gotoProjectPg(usr *User, w http.ResponseWriter, r *http.Request) {

	projects, err := fetchProjectsForUserId(usr.Userid)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(projects) > 0 {
		var data = struct {
			Username string
			Userid   int
			Projects []Project
		}{usr.Name, usr.Userid, projects}
		parsedTemplate.ExecuteTemplate(w, "project_list.html", data)

	} else {
		create_projectPg(usr, w, r)
	}

}

func newProject(w http.ResponseWriter, r *http.Request) {
	useridstr := r.URL.Path[len("/project/newprj/"):]
	if len(useridstr) > 0 {
		numpattern, _ := regexp.Compile("[0-9]+$")
		if numpattern.MatchString(useridstr) {
			userid, err := strconv.ParseInt(useridstr, 10, 32)
			if err != nil {
				error := Error{1, err.Error()}
				showError(error, w)
			} else {
				usr, er := fetchUserForId(int(userid))
				if er != nil {
					error := Error{1, er.Error()}
					showError(error, w)
				}
				create_projectPg(usr, w, r)
			}
		}
	}

}

func create_projectPg(usr *User, w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Username string
		Userid   int
	}{usr.Name, usr.Userid}
	err := parsedTemplate.ExecuteTemplate(w, "create_project.html", data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func create_project(w http.ResponseWriter, r *http.Request) {
	var userid int
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form", err)
		error := Error{1, err.Error()}
		showError(error, w)
		return
	}

	name := r.FormValue("name")
	desc := r.FormValue("desc")
	userid, _ = strconv.Atoi(r.FormValue("userid"))

	var project *Project = new(Project)
	project.Name = name
	project.des = desc
	project.UserId = userid
	project.Completed = false

	err = saveorUpdateProject(project)

	if err != nil {
		error := Error{1, err.Error()}
		showError(error, w)

	} else {
		showProjectList(w, r)
	}

}

func showProjectList(w http.ResponseWriter, r *http.Request) {
	var userid int

	userid, _ = strconv.Atoi(r.FormValue("userid"))
	username := r.FormValue("username")

	projects, err := fetchProjectsForUserId(userid)
	if err != nil {
		error := Error{1, err.Error()}
		showError(error, w)
	} else {

		var data = struct {
			Username string
			Userid   int
			Projects []Project
		}{username, userid, projects}
		parsedTemplate.ExecuteTemplate(w, "project_list.html", data)
	}

}

func getProject(w http.ResponseWriter, r *http.Request) {
	projectidstr := r.URL.Path[len("/project/"):]
	if len(projectidstr) > 0 {
		numpattern, _ := regexp.Compile("[0-9]+$")
		if numpattern.MatchString(projectidstr) {
			projectid, err := strconv.ParseInt(projectidstr, 10, 32)
			if err != nil {
				error := Error{1, err.Error()}
				showError(error, w)
			} else {
				showTaskList(int(projectid), w, r)
			}
		}
	} else {
		error := Error{1, "Invalid Project Id"}
		showError(error, w)
	}
}
