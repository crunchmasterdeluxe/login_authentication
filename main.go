/*
	This app accepts a form POST and authenticates the given
	username and password against a SQLite database.

	It also has built-in functionality to add new users and their
	raw passwords to the users table in the database.

	Local Packages
	-----
		All database functions are stored in the database package.

		The helpers package contains an error handling function used
		throughout the program.

	Future Iterations
	-----
		- Hash passwords before storing
		- Build out front-end more
		- Authorize with a jwt token that can be used for
		  all future requests in app
*/

package main

import (
	"fmt"
	"net/http"

	"crunchmasterdeluxe.com/api_authentication/database"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handleLoginPage)
	r.HandleFunc("/api/login", handleLogin)

	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.Handle("/api/", r)
	http.ListenAndServe("localhost:8080", nil)
}

func handleLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Perform authentication logic
	message, isValidUser := database.Controller(username, password, "read")
	if isValidUser {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "%s"}`, message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"message": "%s"}`, message)
	}
}
