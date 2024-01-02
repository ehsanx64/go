package main

import (
	"fmt"
	"log"
	"net/http"
)

func LoggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo ...")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar ...")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "root ...")
	})

	http.HandleFunc("/foo", LoggerMiddleware(foo))
	http.HandleFunc("/bar", LoggerMiddleware(bar))

	http.ListenAndServe(":8080", nil)
}
