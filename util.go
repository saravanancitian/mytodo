package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func showError(error Error, w http.ResponseWriter) {
	parsedTemplate.ExecuteTemplate(w, "error.html", error)
}

func loadProp(fpath string) map[string]string {
	var props map[string]string = make(map[string]string)
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Fatal("Error reading prop")
	}

	datastr := string(data)
	var lines []string = strings.Split(datastr, "\n")
	for _, line := range lines {

		//Remove comment
		hashidx := strings.IndexByte(line, '#')
		if hashidx != -1 {
			line = strings.TrimSpace(string(line[:hashidx]))
		}

		//create prop map
		if len(line) > 0 {
			var key, val string
			equalidx := strings.IndexByte(line, '=')
			if equalidx >= 0 {
				key = strings.TrimSpace(string(line[:equalidx]))
				val = strings.TrimSpace(string(line[equalidx+1:]))
				props[key] = val
			} else {
				fmt.Printf("Error in line: %s ", line)
			}

		}

	}
	return props
}
