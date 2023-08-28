# API Authentication App

This Go application provides API authentication by accepting a form POST request with a username and password, and then authenticating the given credentials against a SQLite database. The application also includes functionality to add new users and their raw passwords to the users table in the database.

## Explainer Videos
For all you visual learners out there, I started a Youtube thing (we haven't DTR'd) where I take you through everything that I learned as I coded this. Feedback is always welcome!
[First iteration (gin-api branch)](https://youtu.be/h08wQNC0o_o)
[Second iteration (mux-api + unity branches)](https://youtu.be/03zb3sGO1I0)

## Local Packages
### database

The database package contains all the necessary functions related to interacting with the SQLite database. This package handles authentication and user management operations.

### helpers

The helpers package includes a utility error handling function used throughout the application to ensure proper error reporting and handling.

## Future Iterations

In future iterations of this application, the following enhancements are planned:

Hashed Passwords: Passwords will be securely hashed before storing them in the database. This will enhance the security of stored user credentials.

JWT Token Authorization: Implement authorization using JSON Web Tokens (JWT). JWT tokens will be used for all future requests within the app, improving security and user experience.

## Usage

Serve the `./off_grid_solar_home_webgl/index.html` page. Make sure that the iFrame source in `./templates/index.html` matches the local port that the Unity scene is running on.

Compile and Run the App: To run the application, compile the main.go file using `go run main.go`. This will start the server on localhost:8080.

Login Form Submission: Submit a POST request to /login with the username and password fields to authenticate the user. The application will perform the authentication logic and respond with appropriate messages based on the authentication outcome.

Creating Users (Temporary): The code includes a temporary block that demonstrates how to create a new user. This is not part of the main functionality and should be removed or properly integrated in future versions.

## Disclaimer

This application is a basic demonstration of API authentication and is not production-ready. It lacks important security features like hashed passwords and proper front-end design. Use it as a learning resource and starting point for further development.
