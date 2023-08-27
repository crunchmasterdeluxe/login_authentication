package database

import (
	"database/sql"

	"crunchmasterdeluxe.com/api_authentication/helpers"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type User struct {
	ID       int
	Username string
	Password string
}

func initDb() {
	// Initialize a SQLite db and create table
	var err error
	db, err = sql.Open("sqlite3", "./user.db")

	helpers.RaiseError(err, "")

	createTable()
}

func createTable() {
	// Create users table if it doesn't exist
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			password TEXT
		)
	`)

	helpers.RaiseError(err, "")

	_, err = statement.Exec()

	helpers.RaiseError(err, "")

}

func createUser(username string, password string) string {
	// Create and save a new user
	var newUser User = User{
		Username: username,
		Password: password,
	}
	userMessage, err := saveUser(newUser)
	if err != nil {
		helpers.RaiseError(err, "Failed to save user")
	}
	return userMessage
}

func saveUser(user User) (string, error) {
	// Save the provided user to the users table
	// SQL injection note: Appending ; DELETE FROM users WHERE id = 2; doesn't inject, must be abstracted
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	helpers.RaiseError(err, "")
	var userMessage string = "User saved successfully!"
	return userMessage, nil
}

func isValidUser(username, password string) (string, bool) {
	// Check if 1)username exists, and 2)password provided matches password stored.
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		return "User not found", false
	}

	if password != storedPassword {
		return "Invalid credentials", false
	}

	return "Login successful!", true
}

func Controller(username string, password string, queryType string) (string, bool) {
	/*
		Run functions to create a new user in db or authenitcate user

		Process:
			- InitDb (public)
				- createTable
			|
			- createUser
				- saveUser
			|
			- isValidUser
	*/
	initDb()
	if queryType == "create" {
		var message string = createUser(username, password)
		return message, true
	} else if queryType == "read" {
		message, success := isValidUser(username, password)
		return message, success
	}
	return "Invalid Database Procedure. Valid args include `create` or `read`", false
}
