package main

import (
    "log"
    "net/http"
		"html/template"
		"fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "POST done")
	}
	var data = ""
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}