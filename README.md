# Authentication Module

The **Authentication Module** is a simple web application built with **Go**, designed to handle user authentication tasks efficiently. It provides essential functionalities such as:

- **User Registration**
- **User Login**
- **Retrieving User Information**

This project utilizes the **Echo** framework for HTTP handling, **PostgreSQL** for database operations via **GORM**, and **JSON Web Tokens (JWT)** for stateless authentication. It serves as an educational resource or a foundation for more complex authentication systems.

## Technologies Used

- **Go** (Golang)
- **Echo** (Web Framework)
- **PostgreSQL** (Database)
- **GORM** (ORM for Go)
- **JWT** (Authentication)

## Project Structure

```
📂 authentication-module/
├── 📄 main.go            # Entry point, sets up Echo server and routes
├── 📂 handlers/          # Contains HTTP handlers for signup, login, and user retrieval
├── 📂 models/            # Defines the User model with password hashing and verification methods
├── 📂 middleware/        # Includes JWT authentication middleware
└── 📂 db/                # Manages database connections and migrations
```

## Installation & Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/authentication-module.git
   cd authentication-module
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up the PostgreSQL database and configure the connection settings in `.env`.

4. Run database migrations:
   ```sh
   go run db/migrate.go
   ```

5. Start the server:
   ```sh
   go run main.go
   ```

## API Endpoints

- **POST /signup** - Register a new user
- **POST /login** - Authenticate and receive a JWT token
- **GET /user** - Retrieve user information (requires authentication)
- **GET /protected** - for testing JWT token (requires authentication)



---

Feel free to contribute and enhance the module!

