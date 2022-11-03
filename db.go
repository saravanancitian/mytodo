package main

import (
	"database/sql"
)

var db *sql.DB
var insert_user, insert_project, insert_task *sql.Stmt
var update_password, update_project, update_task *sql.Stmt

func openDb() {
	var err error
	db, err = sql.Open("mysql", "root:root@/mytodo")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)
	insert_user, err = db.Prepare("insert into user( name, email, password) values(?,?,?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	insert_project, err = db.Prepare("insert into project(name, userid, completed, des) values(?,?,?, ?)")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	insert_task, err = db.Prepare("insert into task(name, des, projectid, state) values(?,?,?,?)")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	update_password, err = db.Prepare("update user set password = ? where email = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	update_project, err = db.Prepare("update project set name = ?, userid = ?, completed = ?, des = ? where projectid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	update_task, err = db.Prepare("update task set name = ?, des= ?,projectid = ?, state = ? where taskid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func fetchUserForEmail(email string) (*User, error) {
	var usr User
	rows := db.QueryRow("select * from user where email=?", email)
	err := rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func fetchUserForId(userid int) (*User, error) {
	var usr User
	rows := db.QueryRow("select * from user where userid=?", userid)
	err := rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func saveUser(usr *User) error {
	_, err := insert_user.Exec(usr.Name, usr.Email, usr.Password)
	return err
}

func fetchProjectForId(projectid int) (*Project, error) {
	var project Project
	row := db.QueryRow("select * from project where projectid = ?", projectid)

	err := row.Scan(&project.ProjectId, &project.Name, &project.UserId, &project.Completed, &project.des)

	if err != nil && row.Err() == err {
		return nil, err
	}
	return &project, nil
}

func saveorUpdateProject(project *Project) error {
	var err error = nil
	if project.ProjectId == 0 {
		_, err = insert_project.Exec(project.Name, project.UserId, project.Completed, project.des)
	} else {
		_, err = update_project.Exec(project.Name, project.UserId, project.Completed, project.des, project.ProjectId)
	}

	return err
}

func fetchProjectsForUserId(userid int) ([]Project, error) {
	var projects []Project
	rows, err := db.Query("Select * from project where userid = ?", userid)
	if err != nil {
		return nil, err
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

	return projects, nil
}

func fetchTasksForProjectId(projectId int) ([]Task, error) {
	var tasks []Task
	rows, err := db.Query("Select * from task where projectid = ?", projectId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task

		err = rows.Scan(&task.TaskId, &task.Name, &task.Des, &task.ProjectId, &task.State)
		if err != nil {

			break
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}

func fetchTaskforId(taskid int) (*Task, error) {
	var task Task
	row := db.QueryRow("select * from task where taskid = ?", taskid)

	err := row.Scan(&task.TaskId, &task.Name, &task.Des, &task.ProjectId, &task.State)

	if err != nil && row.Err() == err {
		return nil, err
	}
	return &task, nil
}

func saveOrUpdateTask(task *Task) error {
	var err error = nil
	if task.TaskId == 0 {
		_, err = insert_task.Exec(task.Name, task.Des, task.ProjectId, task.State)
	} else {
		_, err = update_task.Exec(task.Name, task.Des, task.ProjectId, task.State, task.TaskId)
	}

	return err
}

func closeDb() {
	insert_user.Close()
	insert_project.Close()
	insert_task.Close()
	update_password.Close()
	update_project.Close()
	update_task.Close()
	db.Close()
}
