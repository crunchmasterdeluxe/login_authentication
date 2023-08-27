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

// every executable package in Go needs to have package called main

// go mod init api_authentication.
// Unused imports are compile time errors
import (
	"fmt"
	"net/http"

	"crunchmasterdeluxe.com/api_authentication/database"

	"github.com/gin-gonic/gin"
)

func main() {

	if 1 == 0 {
		userMessage, success := database.Controller("person6", "secretpassword6", "create")
		fmt.Println(userMessage, success)
	}

	router := gin.Default()

	// Render the login form -
	// TODO could not get to work
	// router.GET("/login-form", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })

	// Handle the login form submission
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Perform authentication logic
		message, isValidUser := database.Controller(username, password, "read")
		if isValidUser {
			c.JSON(http.StatusOK, gin.H{"message": message})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": message})
		}
	})

	router.Run("localhost:8080")
}
