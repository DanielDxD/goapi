# Go Clean Architecture Learning Project

This project is a REST API built with Go, created to practice and understand Clean Architecture principles and modern Go standard library features.

## 🚀 Overview

The main goal of this application is to explore how to properly structure a Go backend application by separating concerns into distinct layers. It implements a simple user management API with in-memory storage, allowing the extraction of core concepts without the overhead of database configurations.

## 🧠 Concepts Learned

During the development of this project, several key software engineering concepts and Go-specific features were applied:

### 1. Clean Architecture (`internal/` directory structure)
The project strictly enforces separation of concerns by dividing the codebase into three main layers:
- **Handlers (Presentation Layer)**: Located in `internal/handlers`. Responsible for handling HTTP requests, decoding JSON payloads, and returning HTTP responses. It has no business logic.
- **Use Cases (Business Logic Layer)**: Located in `internal/usecases`. Contains the core application rules (e.g., verifying if an email already exists before creating a user).
- **Repositories (Data Access Layer)**: Located in `internal/repositories`. Responsible for storing and retrieving data. Currently implemented using an in-memory slice of users for simplicity.

### 2. Dependency Injection
Dependencies are injected explicitly via constructors (`New()` functions). 
- The `Handlers` constructor receives a pointer to `UseCases`.
- The `UseCases` constructor receives a pointer to `Repositories`.

This practice makes the code highly modular, easily testable, and loosely coupled.

### 3. Modern Go Routing (Go 1.22+)
The project utilizes the enhanced routing capabilities introduced in Go 1.22's standard library (`net/http`). It uses pattern matching with HTTP methods directly in `http.HandleFunc`, such as:

```go
http.HandleFunc("GET /users", h.getAllUsers)
http.HandleFunc("POST /users", h.addUser)
```
This eliminates the need for third-party routing libraries (like `gin`, `echo`, or `chi`) for simple REST APIs, keeping dependency count to a minimum.

### 4. Structured Logging (`log/slog`)
The project embraces the `log/slog` package from the standard library to provide leveled and structured logging (e.g., `slog.Info`, `slog.Error`). This makes logs highly readable and easy to parse in production environments.

### 5. Standard JSON Serialization
Handling request and response payloads using standard library `encoding/json` mechanisms:
- Decode incoming request bodies with `json.NewDecoder()`.
- Encode outgoing JSON responses with `json.NewEncoder()`.

## 🛠️ API Endpoints

### 1. Get All Users
- **Method & Route:** `GET /users`
- **Description:** Returns a list of all registered users.
- **Response Structure:** Array of `User` objects.

### 2. Create User
- **Method & Route:** `POST /users`
- **Description:** Creates a new user. Contains business logic that prevents duplicate inserts by checking if the informed email already exists.
- **Request Body Payload Example:**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

## 💻 Running the Project

```bash
# Assuming you are in the project root directory
go run cmd/api/main.go
```
The server will start locally and listen on `http://localhost:3030`. You can test the endpoints using Postman, cURL, or any other HTTP client.
