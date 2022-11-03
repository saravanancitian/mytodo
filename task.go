package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Task struct {
	TaskId    int
	Name      string
	Des       string
	ProjectId int
	State     string
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskidstr := r.URL.Path[len("/task/"):]
	if len(taskidstr) > 0 {
		numpattern, _ := regexp.Compile("[0-9]+$")
		if numpattern.MatchString(taskidstr) {
			taskid, err := strconv.ParseInt(taskidstr, 10, 32)
			if err != nil {
				error := Error{1, err.Error()}
				showError(error, w)
			} else {
				var task *Task
				task, err = fetchTaskforId(int(taskid))
				if err != nil {
					error := Error{1, err.Error()}
					showError(error, w)
				} else {
					parsedTemplate.ExecuteTemplate(w, "add_update_task.html", task)
				}
			}
		}

	} else {
		error := Error{1, "Invalid task Id"}
		showError(error, w)
	}
}

func saveTask(w http.ResponseWriter, r *http.Request) {
	var err error
	taskidstr := r.URL.Path[len("/task/save/"):]
	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form", err)
		error := Error{1, err.Error()}
		showError(error, w)
		return
	}

	name := r.FormValue("name")
	desc := r.FormValue("desc")
	projectid, _ := strconv.Atoi(r.FormValue("projectid"))

	var taskid int64 = 0

	if len(taskidstr) > 0 {
		numpattern, _ := regexp.Compile("[0-9]+$")
		if numpattern.MatchString(taskidstr) {
			taskid, err = strconv.ParseInt(taskidstr, 10, 32)
			if err != nil {
				error := Error{1, err.Error()}
				showError(error, w)
				return
			}
		}
	}

	if taskid > 0 {
		task, _ := fetchTaskforId(int(taskid))
		task.Name = name
		task.Des = desc
		//TODO: state
		saveOrUpdateTask(task)
	} else {
		// new task
		var task *Task = new(Task)
		task.Name = name
		task.Des = desc
		task.ProjectId = projectid
		task.State = "new"

		saveOrUpdateTask(task)
	}
	showTaskList(projectid, w, r)
}

func newTask(w http.ResponseWriter, r *http.Request) {
	projectidstr := r.URL.Path[len("/task/newtask/"):]
	numpattern, _ := regexp.Compile("[0-9]+$")
	if numpattern.MatchString(projectidstr) {
		projectid, err := strconv.ParseInt(projectidstr, 10, 32)
		if err != nil {
			error := Error{1, err.Error()}
			showError(error, w)
		} else {
			var task *Task = new(Task)
			task.ProjectId = int(projectid)
			parsedTemplate.ExecuteTemplate(w, "add_update_task.html", task)
		}
	}
}

func showTaskList(projectid int, w http.ResponseWriter, r *http.Request) {
	tasks, err := fetchTasksForProjectId(projectid)
	if err != nil {
		error := Error{1, err.Error()}
		showError(error, w)
	} else {
		project, err := fetchProjectForId(projectid)
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
					ProjectId   int
					ProjectName string
					ProjectDesc string
					Tasks       []Task
				}{user.Name, project.ProjectId, project.Name, project.des, tasks}
				parsedTemplate.ExecuteTemplate(w, "task_list.html", data)
			}
		}

	}
}
