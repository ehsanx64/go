package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "Todo List",
			Todos: []Todo{
				{
					Title: "Item 1",
					Done:  true,
				},
				{
					Title: "Item 2",
					Done:  true,
				},
				{
					Title: "Item 3",
					Done:  false,
				},
				{
					Title: "Item 4",
					Done:  true,
				},
			},
		}

		tpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
