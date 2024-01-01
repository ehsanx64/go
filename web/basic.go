package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "9999"

/*
** Get remote IP of the requesting client
 */
func getIp(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	} else {
		return r.RemoteAddr
	}
}

/*
** Handler for the '/' route
 */
func handler(w http.ResponseWriter, r *http.Request) {
	// Get the second part of the url
	aname := r.URL.Path[1:]

	// If nothing's there use Go as the default
	if aname == "" {
		aname = "Go"
	}

	// Get the user agent and remote ip
	userAgent := r.UserAgent()
	userIp := getIp(r)

	fmt.Fprintf(w, "Hi there, I love %s!\nUser agent: %s\n", aname, userAgent)
	log.Printf("*** New request:\n")
	log.Printf("%-12s: %s\n", "User Agent", userAgent)
	log.Printf("%-12s: %s\n", "User IP", userIp)
}

func main() {
	log.Println("Server listenning on port " + port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
