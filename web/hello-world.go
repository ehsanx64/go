package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write to response
		fmt.Fprintf(w, "Hello, your request was: "+r.URL.Path)

		// Log to console
		log.Println("*** New request")
		log.Println(r)
	})

	http.ListenAndServe(":8080", nil)
}
