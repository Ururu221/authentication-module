Project Overview
The authentication module is a simple web application built with Go, designed to handle user authentication tasks. It provides functionalities for user registration, login, and retrieving user information, leveraging the Echo framework for HTTP handling, PostgreSQL for database operations via GORM, and JSON Web Tokens (JWT) for stateless authentication. This module is suitable for educational purposes or as a starting point for more complex authentication systems.

The project structure includes:

main.go: The entry point, setting up the Echo server and routes.
handlers/: Contains HTTP handlers for signup, login, and user information retrieval.
models/: Defines the User model with methods for password hashing and verification.
middleware/: Includes JWT authentication middleware.
db/: Manages database connections and migrations.
