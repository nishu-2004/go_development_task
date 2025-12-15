
# User Management API (Go + Fiber)

A RESTful backend API built with **Go** and **Fiber** to manage users using their date of birth (DOB) and dynamically calculate age.  
The project follows clean architecture principles and uses PostgreSQL with SQLC for type-safe database access.

---

## 🚀 Tech Stack

- Go (Fiber)
- PostgreSQL
- SQLC
- pgx
- Uber Zap (logging)
- go-playground/validator

---

## 📁 Project Structure

```

cmd/server/main.go
internal/
handler/
service/
repository/
routes/
middleware/
logger/
db/
db/
sqlc/
sqlc.yaml

````

---

## 🛠️ Setup Instructions

### 1️⃣ Prerequisites

- Go 1.22+
- PostgreSQL (running locally)
- SQLC installed

---

### 2️⃣ Create Database

```sql
CREATE DATABASE users_db;
````

---

### 3️⃣ Create Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

### 4️⃣ Configure Database Connection

Update the connection string in:

```
internal/db/postgres.go
```

Example:

```go
postgres://postgres:<your_password>@localhost:5432/users_db?sslmode=disable
```

---

### 5️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 6️⃣ Generate SQLC Code

```bash
sqlc generate
```

---

### 7️⃣ Run the Server

```bash
go run ./cmd/server
```

Server will start at:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### ➕ Create User

**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### 🔍 Get User by ID

**GET** `/users/{id}`

---

### 📋 List All Users

**GET** `/users`

---

## ✅ Validation Rules

* Name is required and must be at least 2 characters
* DOB must be in `YYYY-MM-DD` format
* DOB cannot be in the future

---

## 🧠 Design Notes

* Age is calculated dynamically using Go’s `time` package
* DOB is stored in the database; age is not stored
* SQLC provides type-safe and performant database access
* Clean separation of concerns using handler → service → repository layers
* Structured logging using Uber Zap
* Request ID and request duration middleware implemented

---

## 📦 Future Enhancements

* Pagination for user listing
* Docker support
* Unit tests
* CI/CD pipeline

---

## 👤 Author

Nishanth P Kashyap
