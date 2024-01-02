package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}

func main() {
	tpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Message: r.FormValue("message"),
		}

		_ = details

		log.Println("*** Contact form submission...")
		log.Println("Name:    ", details.Name)
		log.Println("Email:   ", details.Email)
		log.Println("Message: ", details.Message)

		tpl.Execute(w, struct {
			Success bool
		}{
			true,
		})
	})

	http.ListenAndServe(":8080", nil)
}
