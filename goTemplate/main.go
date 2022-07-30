package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Task string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	templ := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "TODO List",
			Todos: []Todo{
				{Task: "Clean house", Done: false},
				{Task: "Walk pet", Done: false},
				{Task: "Cook meal", Done: false},
				{Task: "Study", Done: false},
			},
		}
		templ.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)

}
