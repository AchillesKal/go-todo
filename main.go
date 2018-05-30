package main

import (
    "log"
    "net/http"
		"html/template"
		"fmt"
		"database/sql"
		"strconv"

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
	database, _ := sql.Open("sqlite3", "./app.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (task, body) VALUES (?, ?)")
	statement.Exec("Task1", "Lorem ipsum")
	rows, _ := database.Query("SELECT id, task, body FROM tasks")
	var id int
	var task string
	var body string
	for rows.Next() {
			rows.Scan(&id, &task, &body)
			fmt.Println(strconv.Itoa(id) + ": " + task + " " + body)
	}

	http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}