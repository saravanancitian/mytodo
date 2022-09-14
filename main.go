package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	userid   int
	name     string
	email    string
	password string
}

var db *sql.DB
var insert_user *sql.Stmt

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
		rows.Scan(&usr.userid, &usr.name, &usr.email, &usr.password)
		fmt.Printf("%s %s %s", usr.name, usr.email, usr.password)
	}

}

func main() {
	fmt.Printf("test")
	openDb()
	testdb()
	defer insert_user.Close()

	defer db.Close()

}
