package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
** Dump the query string
 */
func dumpQuery(r *http.Request) {
	log.Println("*** Query string dump")

	if len(r.URL.Query()) > 0 {
		// If map is not empty loop over its keys/values
		for k, v := range r.URL.Query() {
			log.Printf("%-5s: %s\n", k, v)
		}
	} else {
		// If map is empty log it
		log.Println("Query string is empty")
	}
}

func main() {
	// Add root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request (" + r.URL.Path + ") from " + r.RemoteAddr)

		fmt.Fprintf(w, `
			<a href="http://localhost:8080/?name=adam&age=30">Query</a>
		`)

		// Try to dump all query string parameters
		dumpQuery(r)

		// Get a specific parameter
		log.Println("Name:", r.URL.Query().Get("name"))
	})

	// Add static file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
