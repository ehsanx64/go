package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Add root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome ...")
		fmt.Println("New request (" + r.URL.Path + ") from " + r.RemoteAddr)
	})

	// Add static file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
