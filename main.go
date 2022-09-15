package main

import (
	"database/sql"
	"fmt"
	"net/http"

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

func closeDb() {
	insert_user.Close()
	db.Close()
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

func login(w http.ResponseWriter, r *http.Request) {
	var email string
	var password string
	var qpass string
	email = r.PostForm.Get("email")
	password = r.PostForm.Get("password")
	fmt.Println(email + " " + password)
	rows := db.QueryRow("select password from user where email=?", email)

	rows.Scan(&qpass)
	if password == qpass {
		w.Write([]byte("Success"))

	} else {
		w.Write([]byte("Failure"))
	}

}

func main() {
	openDb()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":3000", nil)
	defer closeDb()
}
