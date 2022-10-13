package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	openDb()
	defer closeDb()
	iniweb()
	http.ListenAndServe(":3000", servermux)
}
