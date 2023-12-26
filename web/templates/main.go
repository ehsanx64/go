package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func generateTodos() []Todo {
	var res []Todo

	for i := 1; i < 10; i++ {
		res = append(res, Todo{
			Title: "Item " + strconv.Itoa(i),
			Done:  generateRandomBool(),
		})
	}

	return res
}

func generateRandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func main() {
	tpl := template.Must(template.ParseFiles("layout.html", "todo-list.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "Todo List",
			Todos:     generateTodos(),
		}

		tpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
