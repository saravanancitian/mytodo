package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	openDb()
	iniweb()

	http.HandleFunc("/js/", servJs)

	http.HandleFunc("/", servhtml)
	http.HandleFunc("/*.html", servhtml)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", nil)
	defer closeDb()
}
