package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	cookieName = "sess"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

/*
** This route is ok only for authenticated users
 */
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// If it isn't forbidden for the user
	fmt.Fprintln(w, "A cake for you!")
}

/*
** User would be authenticated within this route
 */
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	// Authetication code
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	// Revoke user's authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
