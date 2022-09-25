package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Userid   int    `json:"userid,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
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

	http.Handle("/js", http.FileServer(http.Dir("static/js/")))

	http.HandleFunc("/", indexhtml)
	http.HandleFunc("/signin.html", signinhtml)
	http.HandleFunc("/signup.html", signuphtml)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", nil)
	defer closeDb()
}
