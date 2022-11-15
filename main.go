package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	props := loadProp("conf/app.prop")
	openDb(props["db.driver"], props["db.datasource"])
	defer closeDb()
	iniweb()
	fmt.Print(" Staring server ")
	http.ListenAndServe(":3000", servermux)

}
