package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var insert_user, insert_project, insert_task *sql.Stmt
var update_password, update_project_status, update_task *sql.Stmt

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

	update_project_status, err = db.Prepare("update project set completed = ? where projectid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	update_task, err = db.Prepare("update task set name = ?, des= ?, state = ? where taskid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func testdb() {
	// result, err := insert_user.Exec("testusr", "testuser@test.com", "password")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result.LastInsertId())

	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr user
		rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)
		fmt.Printf("%s %s %s", usr.Name, usr.Email, usr.Password)
	}

}

func closeDb() {
	insert_user.Close()
	insert_project.Close()
	insert_task.Close()
	update_password.Close()
	update_project_status.Close()
	update_task.Close()
	db.Close()
}
