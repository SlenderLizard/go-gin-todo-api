## Personal Todo API

A secure RESTful Backend API build with GO(1.25.5) and Gin, implements JWT authentication and in-memory data management. In-memory could be changed with db implementation later.

## Packages

* **Gin Gonic** ([gin-gonic.com](https://gin-gonic.com/)): Routing and middleware.
* **JWT (v5)** ([golang-jwt/jwt](https://github.com/golang-jwt/jwt/v5)): Secure user authorization with expiration claims.
* **Bcrypt** ([golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)): Secure password hashing.
* **Godotenv** ([joho/godotenv](https://github.com/joho/godotenv)): Configuration via environment variables.
* **Concurrency**: `sync.RWMutex` for thread-safe in-memory storage.

## Installation
1. Clone the repository
2. Create a .env file in the root directory  
Example:  
```
PORT=3000
JWT_SECRET=your_super_secret_key
```
4. Install dependencies
```
go mod tidy
```
### Running the App  
Execute the following command in the root folder:
go run main.go

## API Endpoints

### Public Routes
* **POST** `/register`: Creates a new user. Username must be unique.
* **POST** `/login`: Authenticate and receive a JWT token.

### Protected Routes
*Requires Header:* `Authorization: Bearer <your_token>`

* **GET** `/todos`: List authenticated user's todos.
* **GET** `/todos/:id`: Get a specific todo item.
* **POST** `/todos`: Create a new todo and link it to the authenticated user.
* **PUT** `/todos/:id`: Update a specific todo (Title, Description, or Completion status).
* **DELETE** `/todos/:id`: Remove a specific todo item.

## Project Structure

* **/handlers**: HTTP request logic and correct HTTP status code management.
* **/repository**: In-memory storage with `sync.RWMutex` for concurrency safety.
* **/middleware**: JWT authentication and todo ownership validation.
* **/models**: Data structures (structs) for User and Todo entities.
* **/initializers**: Environment variables (.env) and initial configuration setup.
* **/postman_collections**: Pre-configured JSON files for easy API testing.

Tried to automate the postman, but failed. Token needs to be pasted manually.
