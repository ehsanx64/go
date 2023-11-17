package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, your request was "+r.URL.Path)
		fmt.Println("New request")
		fmt.Println(r)
	})

	http.ListenAndServe(":8080", nil)
}
