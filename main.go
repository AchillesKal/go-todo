package main

import (
    "log"
    "net/http"
		"html/template"
		"fmt"

		_ "github.com/mattn/go-sqlite3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
	}

	switch r.Method {
	case "GET":
		var data = ""
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, data)
	case "POST":
		var data = ""
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, data)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}