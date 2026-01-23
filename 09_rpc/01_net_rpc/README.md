# net/rpc Example

This example demonstrates Go's built-in RPC (Remote Procedure Call) functionality using the `net/rpc` package. RPC allows you to call methods on remote objects as if they were local.

## Overview

The example implements:
- **RPC Server**: Registers services and handles incoming connections
- **RPC Client**: Makes both synchronous and asynchronous calls to the server
- **Multiple Services**: Arithmetic operations and string operations
- **Error Handling**: Demonstrates proper error handling for RPC calls

## Services

### ArithService
- `Add(a, b int) int` - Returns a + b
- `Multiply(a, b int) int` - Returns a * b
- `Divide(a, b int) float64` - Returns a / b (with division by zero check)
- `Power(a, b int) int` - Returns a^b

### StringService
- `Concat(a, b int) string` - Concatenates string representations of a and b
- `Length(a, b int) int` - Returns length of concatenated string

## Running the Example

```bash
cd golang_roadmap/09_rpc/01_net_rpc
go mod tidy
go run main.go
```

The program will:
1. Start an RPC server on port 1234
2. Run an RPC client that demonstrates various calls
3. Show both synchronous and asynchronous RPC calls
4. Display results and error handling

## Key Concepts Demonstrated

### RPC Method Requirements
- Methods must be exported (start with capital letter)
- Methods must have exactly two arguments
- First argument is the input (any type)
- Second argument is the output (must be a pointer)
- Methods must return an error

### Synchronous Calls
```go
var reply int
err := client.Call("Service.Method", args, &reply)
```

### Asynchronous Calls
```go
call := client.Go("Service.Method", args, &reply, nil)
reply := <-call.Done
```

### Service Registration
```go
service := new(MyService)
rpc.Register(service)
```

## Output Example

```
RPC server starting on port 1234...
Connected to RPC server

=== Synchronous RPC Calls ===
Add(10, 5) = 15
Multiply(10, 5) = 50
Power(10, 5) = 100000
Divide(10, 5) = 2.00
Divide by zero error (expected): division by zero
Concat(10, 5) = 105
Length(10, 5) = 3

=== Asynchronous RPC Calls ===
Async Add(20, 30) = 50
Async Multiply(7, 8) = 56

RPC client finished
```

## Architecture

```
Client Application
        |
        | TCP Connection
        v
RPC Client (net/rpc)
        |
        | Encoded RPC Calls
        v
RPC Server (net/rpc)
        |
        | Method Calls
        v
Registered Services
```

## Advantages of net/rpc

- **Type Safety**: Compile-time type checking
- **Simple API**: Easy to use with Go's built-in types
- **Automatic Serialization**: Handles encoding/decoding automatically
- **Concurrent**: Handles multiple clients simultaneously
- **Standard Library**: No external dependencies

## Limitations

- Only works with Go (not cross-language)
- Requires TCP connections
- No built-in authentication or encryption
- No service discovery

## Resources

- [net/rpc package documentation](https://pkg.go.dev/net/rpc)
- [Introduction to RPC in Go](https://medium.com/@shivambhadani_/introduction-to-rpc-in-go-building-rpc-client-and-server-with-golang-5794675e9a12)