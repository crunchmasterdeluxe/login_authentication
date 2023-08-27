# Database Package

The database package provides functionality for interacting with an SQLite database in the context of the API authentication application.
Package Overview

The database package contains functions for creating and managing user data in the database, including user creation and authentication.

## Table of Contents

- [Initialization](#initialization)
- [Creating Users](#creating-users)
- [Saving Users](#saving-users)
- [Validating Users](#validating-users)
- [Controller Function](#controller-function)

## Initialization

The initDb function initializes an SQLite database and creates the required users table if it doesn't already exist. This function is called when interacting with the database.

## Creating Users

The createUser function is used to create a new user with the provided username and password. It utilizes the saveUser function to save the user details to the database.
## Saving Users

The saveUser function takes a User struct and inserts the user's username and password into the users table in the database. Proper error handling is implemented to ensure successful user data insertion.

## Validating Users

The isValidUser function is responsible for validating user credentials. It checks if the provided username exists in the database and if the provided password matches the stored password for that user.
## Controller Function

The Controller function acts as a central interface for managing user-related operations. It takes a query type as an argument (create or read) and routes the execution to the appropriate sub-functions based on the query type. This function serves as the primary entry point for user authentication and creation.
## Note on Security

The database package does not currently hash passwords before storage, and proper security practices are not fully implemented. In future iterations, passwords should be securely hashed before being stored in the database.

Please note that this package is designed as a demonstration and may lack certain security measures required for production use. It is recommended to further enhance the security aspects before deploying the application.
Usage

To use the functions provided by the database package, you can import it into your Go application:

```go

import (
	"crunchmasterdeluxe.com/api_authentication/database"
)
```

You can then call the functions as needed:

```go

func main() {
	// Example of creating a new user
	userMessage := database.Controller("newuser", "newpassword", "create")
	fmt.Println(userMessage)

	// Example of validating user credentials
	message, success := database.Controller("existinguser", "existingpassword", "read")
	fmt.Println(message, success)
}
```

Please note that the example usage above assumes proper integration into the main application.
