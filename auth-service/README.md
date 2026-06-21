# Auth Service

Authentication service built with Golang, Gin, MySQL, JWT, Docker, and Clean Architecture.

---

# Features

* User Register
* User Login
* JWT Authentication
* Access Token
* Refresh Token
* Refresh Token Rotation
* Logout Session
* Multi Device Session
* Protected Routes
* Swagger Documentation
* Dockerized Service
* MySQL Integration
* Environment Configuration
* Clean Architecture

---

# Tech Stack

* Golang
* Gin Gonic
* MySQL
* JWT
* Docker
* Swagger
* Air Hot Reload

---

# Project Structure

```bash
auth-service/
├── app/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── router/
│   └── domain/
│       ├── dao/
│       └── dto/
├── config/
├── docs/
├── migrations/
├── .env
├── Dockerfile
├── go.mod
└── main.go
```

---

# Architecture Flow

```text
Route
↓
Middleware
↓
Controller
↓
DTO
↓
Service
↓
Repository
↓
DAO
↓
Database
```

---

# Layer Explanation

## Router

Responsible for routing endpoint URLs to controllers.

Example:

```go
r.POST("/login", controller.Login)
```

---

## Middleware

Responsible for filtering requests before entering the controller.

Examples:

* JWT Authentication
* Rate Limiting
* Role Validation

---

## Controller

Responsible for:

* Receiving HTTP requests
* Parsing request body
* Returning JSON responses
* Calling service layer

---

## DTO (Data Transfer Object)

Used for API request and response structures.

Example:

```go
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

---

## Service

Contains business logic.

Examples:

* Password validation
* JWT generation
* Refresh token rotation
* Session management

---

## Repository

Responsible for database communication.

Examples:

* Insert data
* Select data
* Update data
* Delete data

---

## DAO (Data Access Object)

Represents database entities.

Example:

```go
type User struct {
    ID       int
    Name     string
    Email    string
    Password string
}
```

---

# Environment Variables

Create `.env` file:

```env
DB_HOST=mysql-auth
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=auth_db
JWT_SECRET=SECRET_KEY
```

---

# Run with Docker

## Build and Run

```bash
docker compose up --build
```

---

# Run Migration

## Create Migration File

```bash
migrate create -ext sql -dir migrations create_users_table
```

## Run Migration Up

```bash
migrate -path migrations -database "mysql://root:root@tcp(localhost:3308)/auth_db" up
```

## Run Migration Down

```bash
migrate -path migrations -database "mysql://root:root@tcp(localhost:3308)/auth_db" down
```

Example migration:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

# API Endpoints

## Auth

| Method | Endpoint  | Description          |
| ------ | --------- | -------------------- |
| POST   | /register | Register User        |
| POST   | /login    | Login User           |
| POST   | /refresh  | Refresh Access Token |
| POST   | /logout   | Logout User          |
| GET    | /profile  | Get Current User     |

---

# Login Example

## Request

```json
{
  "email": "admin@gmail.com",
  "password": "123456",
  "device_name": "MacBook Chrome"
}
```

## Response

```json
{
  "data": {
    "access_token": "TOKEN",
    "refresh_token": "TOKEN"
  }
}
```

---

# Protected Route Example

Add Authorization Header:

```text
Authorization: Bearer ACCESS_TOKEN
```

---

# Refresh Token Rotation

Flow:

```text
Login
↓
Access Token + Refresh Token
↓
Refresh Token Used
↓
Old Refresh Token Deleted
↓
New Refresh Token Generated
```

---

# Session Management

Supports:

* Multi Device Login
* Logout Current Session
* Logout All Sessions
* Device Tracking
* IP Tracking

---

# Swagger Documentation

Swagger Endpoint:

```text
http://localhost:8081/swagger/index.html
```

Generate Swagger:

```bash
swag init
```

---

# Future Improvements

* Forgot Password
* Email Verification
* RBAC
* Rate Limiting
* Redis Blacklist
* OAuth Login
* Two Factor Authentication

---

# Author

Syeila Ruthby
