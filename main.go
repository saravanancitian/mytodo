package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Userid   int    `json:"userid,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
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
		rows.Scan(&usr.Userid, &usr.Name, &usr.Email, &usr.Password)
		fmt.Printf("%s %s %s", usr.Name, usr.Email, usr.Password)
	}

}

// func testform(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	usr := user{}

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%T", body)
// 	err = json.Unmarshal(body, &usr)
// 	if err != nil {
// 		w.Write([]byte("failed"))
// 	}
// 	fmt.Print(" user " + usr.Email)
// 	w.Write([]byte("success"))
// }

func main() {
	openDb()
	iniweb()

	http.HandleFunc("/", indexhtml)
	http.HandleFunc("/signin.html", signinhtml)
	http.HandleFunc("/signup.html", signuphtml)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", nil)
	defer closeDb()
}
