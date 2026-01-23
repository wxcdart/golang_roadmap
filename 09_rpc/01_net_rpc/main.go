package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Args represents the arguments for RPC calls
type Args struct {
	A, B int
}

// ArithService provides arithmetic operations
type ArithService struct{}

// Add performs addition
func (a *ArithService) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// Multiply performs multiplication
func (a *ArithService) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide performs division with error handling
func (a *ArithService) Divide(args *Args, reply *float64) error {
	if args.B == 0 {
		return fmt.Errorf("division by zero")
	}
	*reply = float64(args.A) / float64(args.B)
	return nil
}

// Power calculates A raised to the power of B
func (a *ArithService) Power(args *Args, reply *int) error {
	result := 1
	for i := 0; i < args.B; i++ {
		result *= args.A
	}
	*reply = result
	return nil
}

// StringService provides string operations
type StringService struct{}

// Concat concatenates two strings
func (s *StringService) Concat(args *Args, reply *string) error {
	// For demo purposes, convert numbers to strings and concatenate
	*reply = fmt.Sprintf("%d%d", args.A, args.B)
	return nil
}

// Length returns the length of a string representation
func (s *StringService) Length(args *Args, reply *int) error {
	str := fmt.Sprintf("%d%d", args.A, args.B)
	*reply = len(str)
	return nil
}

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// Register the services
	arith := new(ArithService)
	stringSvc := new(StringService)

	rpc.Register(arith)
	rpc.Register(stringSvc)

	// Start listening
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	defer listener.Close()

	log.Println("RPC server starting on port 1234...")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		go rpc.ServeConn(conn)
	}
}

func runClient() {
	// Connect to the server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer client.Close()

	log.Println("Connected to RPC server")

	// Test arithmetic operations
	args := &Args{10, 5}
	var reply int

	// Synchronous calls
	fmt.Println("\n=== Synchronous RPC Calls ===")

	// Test Add
	err = client.Call("ArithService.Add", args, &reply)
	if err != nil {
		log.Printf("Add error: %v", err)
	} else {
		fmt.Printf("Add(%d, %d) = %d\n", args.A, args.B, reply)
	}

	// Test Multiply
	err = client.Call("ArithService.Multiply", args, &reply)
	if err != nil {
		log.Printf("Multiply error: %v", err)
	} else {
		fmt.Printf("Multiply(%d, %d) = %d\n", args.A, args.B, reply)
	}

	// Test Power
	err = client.Call("ArithService.Power", args, &reply)
	if err != nil {
		log.Printf("Power error: %v", err)
	} else {
		fmt.Printf("Power(%d, %d) = %d\n", args.A, args.B, reply)
	}

	// Test Divide
	var floatReply float64
	err = client.Call("ArithService.Divide", args, &floatReply)
	if err != nil {
		log.Printf("Divide error: %v", err)
	} else {
		fmt.Printf("Divide(%d, %d) = %.2f\n", args.A, args.B, floatReply)
	}

	// Test division by zero
	zeroArgs := &Args{10, 0}
	err = client.Call("ArithService.Divide", zeroArgs, &floatReply)
	if err != nil {
		fmt.Printf("Divide by zero error (expected): %v\n", err)
	}

	// Test string operations
	var stringReply string
	err = client.Call("StringService.Concat", args, &stringReply)
	if err != nil {
		log.Printf("Concat error: %v", err)
	} else {
		fmt.Printf("Concat(%d, %d) = %s\n", args.A, args.B, stringReply)
	}

	err = client.Call("StringService.Length", args, &reply)
	if err != nil {
		log.Printf("Length error: %v", err)
	} else {
		fmt.Printf("Length(%d, %d) = %d\n", args.A, args.B, reply)
	}

	// Asynchronous calls
	fmt.Println("\n=== Asynchronous RPC Calls ===")

	// Make async calls
	addCall := client.Go("ArithService.Add", &Args{20, 30}, &reply, nil)
	multiplyCall := client.Go("ArithService.Multiply", &Args{7, 8}, &reply, nil)

	// Wait for results
	addReply := <-addCall.Done
	if addReply.Error != nil {
		log.Printf("Async Add error: %v", addReply.Error)
	} else {
		fmt.Printf("Async Add(20, 30) = %d\n", reply)
	}

	multiplyReply := <-multiplyCall.Done
	if multiplyReply.Error != nil {
		log.Printf("Async Multiply error: %v", multiplyReply.Error)
	} else {
		fmt.Printf("Async Multiply(7, 8) = %d\n", reply)
	}

	fmt.Println("\nRPC client finished")
}

func main() {
	var wg sync.WaitGroup

	// Start server in background
	wg.Add(1)
	go startServer(&wg)

	// Wait a bit for server to start
	time.Sleep(100 * time.Millisecond)

	// Run client
	runClient()

	// Wait for server to finish (it won't in this case)
	// wg.Wait()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down...")
}
