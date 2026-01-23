# net/http REST API Example

This folder contains a comprehensive example demonstrating building a REST API using Go's standard library `net/http` package.

The example implements a user management API with proper error handling, validation, and production-ready features.

## Quick Start

```bash
cd golang_roadmap/08_web_development/01_net_http
go mod tidy
go run main.go
```

The server will start on port 8080. Test with curl:

```bash
# Get all users
curl http://localhost:8080/users

# Create a new user
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice"}' http://localhost:8080/users

# Test error cases
curl -X POST -H "Content-Type: application/json" -d '{"name":}' http://localhost:8080/users
curl -X POST -H "Content-Type: text/plain" -d '{"name":"Bob"}' http://localhost:8080/users
curl -X PUT http://localhost:8080/users
```

## Features Demonstrated

- **HTTP Server Setup**: Using `http.Server` with timeouts and `http.ServeMux`
- **REST API Design**: GET and POST endpoints with proper HTTP methods
- **JSON Handling**: Encoding/decoding with `encoding/json`
- **Middleware**: Logging middleware for request tracking
- **Error Handling**: Comprehensive error responses with appropriate HTTP status codes
- **Input Validation**: Content-type checking, JSON validation, required field validation
- **Thread Safety**: Mutex-protected shared state
- **Graceful Shutdown**: Signal handling and server shutdown with timeout
- **HTTP Status Codes**: Proper use of 200, 201, 400, 405, 415 status codes

## API Endpoints

- `GET /users` - Returns list of all users as JSON
- `POST /users` - Creates a new user from JSON payload

## Error Responses

- Invalid JSON: `400 Bad Request` with "Invalid JSON"
- Wrong content-type: `415 Unsupported Media Type` with "Content-Type must be application/json"
- Missing required fields: `400 Bad Request` with "Name is required"
- Invalid HTTP methods: `405 Method Not Allowed`

## Resources

- [net/http package in Go](https://medium.com/@emonemrulhasan35/net-http-package-in-go-e178c67d87f1)
- [How To Make an HTTP Server in Go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)