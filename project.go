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

	err = saveProject(project)

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

	var data = struct {
		Username string
		Projects []Project
	}{username, projects}
	parsedTemplate.ExecuteTemplate(w, "project_list.html", data)

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
				tasks, err := fetchTasksForProjectId(int(projectid))
				if err != nil {
					error := Error{1, err.Error()}
					showError(error, w)
				} else {
					project, err := fetchProjectForId(int(projectid))
					if err != nil {
						error := Error{1, err.Error()}
						showError(error, w)
					} else {
						user, err := fetchUserForId(project.UserId)
						if err != nil {
							error := Error{1, err.Error()}
							showError(error, w)
						} else {
							var data = struct {
								UserName    string
								ProjectName string
								ProjectDesc string
								Tasks       []Task
							}{user.Name, project.Name, project.des, tasks}
							parsedTemplate.ExecuteTemplate(w, "task_list.html", data)
						}
					}

				}
			}
		}
	} else {
		error := Error{1, "Invalid Project Id"}
		showError(error, w)
	}
}
