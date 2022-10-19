package main

import "net/http"

func showError(error Error, w http.ResponseWriter) {
	parsedTemplate.ExecuteTemplate(w, "error.html", error)
}
