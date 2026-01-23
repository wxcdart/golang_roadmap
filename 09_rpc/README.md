# RPC (Remote Procedure Call) Examples

This directory contains examples of Remote Procedure Call implementations in Go.

## 01_net_rpc

Demonstrates Go's built-in `net/rpc` package for building RPC servers and clients.

**Features:**
- RPC server with multiple services
- Synchronous and asynchronous client calls
- Error handling and type safety
- TCP-based communication

**Run:**
```bash
cd 01_net_rpc
go run main.go
```

The example shows arithmetic and string operations being called remotely between a client and server running in the same process.