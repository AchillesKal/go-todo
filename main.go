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

type Task struct {
	Title string
	Body  error
}

func getTasks() string {
	database, _ := sql.Open("sqlite3", "./app.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO tasks (title, body) VALUES (?, ?)")
	statement.Exec("Task1", "Lorem ipsum")
	rows, _ := database.Query("SELECT id, title, body FROM tasks")
	var id int
	var title string
	var body string
	for rows.Next() {
			rows.Scan(&id, &title, &body)
			fmt.Println(strconv.Itoa(id) + ": " + title + " " + body)
	}
	return title
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
	}

	switch r.Method {
	case "GET":
		var data = getTasks()
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, data)
	case "POST":
		var data = getTasks()

		r.ParseForm()
		x := r.Form.Get("title")
		y := r.Form.Get("body")
		fmt.Println(x)
		fmt.Println(y)


		database, _ := sql.Open("sqlite3", "./app.db")

		statement, _ := database.Prepare("INSERT INTO tasks(title, body) values(?,?)")
		statement.Exec(x, y)

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