package main

import (
	"fmt"
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
					fmt.Printf("%s", task.Name)
				}
			}
		}

	} else {
		error := Error{1, "Invalid task Id"}
		showError(error, w)
	}
}
