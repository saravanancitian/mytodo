package main

import (
	"log"
	"net/http"
)

type Project struct {
	ProjectId int
	Name      string
	UserId    int
	Completed bool
	des       string
}

func gotoProjectPg(usr *User, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("Select * from project where userid = ?", usr.Userid)
	if err != nil {
		//TODO need check for error
		log.Fatal("Error in query")
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project
		err = rows.Scan(&project.ProjectId, &project.Name, &project.UserId, &project.Completed, &project.des)
		if err != nil {
			
			break
		}
		projects = append(projects, project)
	}

	if len(projects) > 0 {
		var data = struct {
			username string
			projects []Project
		}{usr.Name, projects}
		parsedTemplate.ExecuteTemplate(w, "project_list.html", data)

	}
	else{
		parsedTemplate.ExecuteTemplate(w, "create_project.html", nil)

	}

}
