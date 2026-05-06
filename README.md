# Go REST API Learning — Day 1

## What I Learned

Today I built a simple REST API in Go while learning backend architecture fundamentals.

---

# Concepts Learned

## 1. HTTP Routing in Go

Learned how to:

* Start a Go HTTP server
* Register routes using `http.HandleFunc`
* Handle requests and responses

Example routes:

* `/ping`
* `/users`
* `/user`

---

# 2. Backend Layered Architecture

Learned how to separate backend responsibilities into layers.

## `handler.go`

Responsible for:

* Handling HTTP requests/responses
* Reading request data
* Validating methods
* Returning JSON/status codes

## `service.go`

Responsible for:

* Business logic
* ID generation
* Application workflows
* Coordinating storage operations

## `storage.go`

Responsible for:

* In-memory database logic
* CRUD operations
* Managing maps/data persistence

## `model.go`

Responsible for:

* Data structures
* Struct definitions
* Representing application data

---

# 3. Structs in Go

Created a `User` model:

Fields:

* ID
* Name
* Email

Learned:

* Capitalized names are exported/public
* Lowercase names are private
* Difference between type names and variable names

Example understanding:

* `User` = type/blueprint
* `user` = variable instance

---

# 4. In-Memory Storage

Built a temporary database using:

```go
map[int]User
```

Concept:

```txt
ID → User
```

Learned:

* Why maps are useful for O(1) lookups
* How Go maps return:

  * value
  * existence boolean

---

# 5. JSON Request Deserialization

Learned how client JSON becomes a Go struct.

Flow:

```txt
Client JSON
    ↓
Request Body
    ↓
json.Decode()
    ↓
Go Struct
```

Key concept:

```go
json.NewDecoder(r.Body).Decode(&req)
```

Learned:

* `r.Body` contains incoming request data
* `Decode()` fills struct fields
* `&req` passes memory address

---

# 6. JSON Response Serialization

Learned how Go structs become JSON responses.

Flow:

```txt
Go Struct
    ↓
json.Encode()
    ↓
JSON Response
```

Used:

```go
json.NewEncoder(w).Encode(user)
```

---

# 7. HTTP Headers

Learned why APIs should send:

```go
w.Header().Set("Content-Type", "application/json")
```

Purpose:

* Tells client response format
* Helps frontend/API clients parse correctly

---

# 8. Error Handling in Go

Learned Go's explicit error handling pattern:

```go
if err != nil
```

Meaning:

* `nil` → no error
* non-nil → error exists

---

# 9. CRUD Operations

Implemented:

## Create User

* POST request
* Decode JSON
* Generate ID
* Store user

## Get User

* GET request
* Query parameter parsing
* String → int conversion
* User lookup

## Delete User

* DELETE request
* Existence check
* Remove from map

## Update User

* PUT request
* Decode updated JSON
* Replace existing map value

---

# 10. Query Parameters

Learned how to read URL query params.

Example:

```txt
/user?id=1
```

Used:

```go
r.URL.Query().Get("id")
```

---

# 11. String to Integer Conversion

Used:

```go
strconv.Atoi()
```

Learned:

* Query params arrive as strings
* Need conversion for map lookups

---

# 12. Multiple Return Values in Go

Learned that functions can return multiple values.

Example:

```go
user, exists := service.GetUser(userID)
```

Used for:

* Map lookups
* Error handling
* Existence checks

---

# 13. Request Lifecycle Understanding

Complete backend flow learned:

```txt
Client Request
      ↓
Handler
      ↓
Service
      ↓
Storage
      ↓
Response Serialization
      ↓
Client Response
```

---

# Important Go Concepts Learned

## `:=`

Used for:

* declaring + initializing variables

## `=`

Used for:

* updating existing variables

---

# What I Built

A mini REST API with:

* Create User
* Get User
* Update User
* Delete User
* Ping Route

Using:

* Go standard library only
* In-memory storage
* Layered architecture
* JSON APIs

---

# Next Things To Learn

* Mutex & concurrency safety
* Route grouping
* Better router libraries
* Middleware
* Validation
* Persistent databases
* Authentication
* REST API best practices
* Goroutines & synchronization
