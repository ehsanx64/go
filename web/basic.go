package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "9999"

func getIp(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	} else {
		return r.RemoteAddr
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	userIp := getIp(r)
	fmt.Fprintf(w, "Hi there, I love %s!\nUser agent: %s\n", r.URL.Path[1:], userAgent)
	fmt.Printf("New request:\n")
	fmt.Printf("User Agent: \t%s\n", userAgent)
	fmt.Printf("User IP: \t%s\n", userIp)
}

func main() {
	log.Println("Server listenning on port " + port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
