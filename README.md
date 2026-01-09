Personal Todo API

A secure RESTful Backend API build with GO(1.25.5) and Gin, implements JWT authentication and in-memory data management. In-memory could .be changed with db implementation later.

Packages

Gin Gonic (https://gin-gonic.com/): Routing and middleware.
JWT (v5) (https://pkg.go.dev/github.com/golang-jwt/jwt/v5):  Secure user authorization with expiration claims.
Bcryipt (https://pkg.go.dev/golang.org/x/crypto/bcrypt): Secure password hashing. 
Godotenv (https://github.com/joho/godotenv): Configuration via enviroment variables.
Concurrency: sync.RWMutex for thread safe in-memory storage.

Installation
1. Clone the repository
2. Create a .env file in the root directory
Example:
PORT=3000
JWT_SECRET=your_super_secret_key
3. Install dependencies
go mod tidy

Running the App
Execute the following command in the root folder:
go run main.go

API Endpoints

Public Routes
POST /register: Creates a new user, username must be unique.
POST /login:  Authenticate and recieve a JWT token.
Protected Routes (Requires Authorization: Bearer <token>)
GET /todos: List authenticated users todos.
GET /todos/:id; Get a specific todo
POST /todos: Create a new todo.
PUT /todos/:id; Update a todo.
DELETE /todos/:id; Remove a specific todo item. 

Project Structure 
/handlers: HTTP request logic and status code management. 
/repository: In-memory storage with mutex implementation.
/middleware: JWT authenticaion and ownership validation.
/models: Structs for User and Todo
/initializers: Setup for .env variables
/postman_collections: Configured JSON files for API testing.

Tried to automate the postman, but failed. Token needs to be pasted manually.