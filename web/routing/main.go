package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Getting the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Adding the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Updating the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Deleting the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// Books handlers (http verbs)
	r.HandleFunc("/books/{title}", GetBook).Methods("GET")
	r.HandleFunc("/books/{title}", AddBook).Methods("POST")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
