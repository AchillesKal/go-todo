package main

import (
    "log"
    "net/http"
		"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var data = ""
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}