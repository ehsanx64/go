package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func encode(w http.ResponseWriter, r *http.Request) {
	user := User{
		FirstName: "Adam",
		LastName:  "Smith",
		Age:       40,
	}

	json.NewEncoder(w).Encode(user)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s %s is %d years old!\n",
		user.FirstName, user.LastName, user.Age)
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	http.ListenAndServe(":8080", nil)
}
