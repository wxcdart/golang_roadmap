// context_example.go
// Demonstrates key features of Go's context package: cancellation, timeouts, deadlines, and value propagation.
//
// About context:
// - context.Context is used to carry deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.
// - Cancellation: Allows you to signal goroutines to stop work early (e.g., user aborts, timeout, parent process exits).
// - Timeout/Deadline: Automatically cancels context after a set duration or at a specific time.
// - Value: Lets you attach key-value pairs to a context (for request-scoped data, e.g., user ID, trace ID).
// - Always pass context as the first argument to functions that need it (by convention).

package main

import (
	"context"
	"fmt"
	"time"
)

// worker simulates a long-running task that checks for context cancellation.
// It listens for ctx.Done() to know when to stop early.
// ctx.Err() gives the reason for cancellation (context.Canceled or context.DeadlineExceeded).
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelled (%v)\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	// Example 1: Manual cancellation
	// context.WithCancel returns a derived context and a cancel function.
	// Call cancel() to signal cancellation to all goroutines using this context.
	ctx1, cancel1 := context.WithCancel(context.Background())
	go worker(ctx1, 1)
	time.Sleep(1 * time.Second)
	fmt.Println("Main: cancelling ctx1")
	cancel1() // Signal cancellation
	time.Sleep(400 * time.Millisecond)

	// Example 2: Timeout
	// context.WithTimeout returns a context that is automatically cancelled after the given duration.
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2() // Always call cancel to release resources
	go worker(ctx2, 2)
	time.Sleep(2 * time.Second)

	// Example 3: Deadline
	// context.WithDeadline cancels the context at a specific time.
	deadline := time.Now().Add(1 * time.Second)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()
	go worker(ctx3, 3)
	time.Sleep(2 * time.Second)

	// Example 4: Value propagation
	// context.WithValue attaches a key-value pair to the context.
	// Use for request-scoped data (not for passing large data or for global state).
	ctx4 := context.WithValue(context.Background(), "userID", 42)
	printValue(ctx4)
}

// printValue demonstrates retrieving a value from context.
// Use context.Value(key) to get a value attached to the context.
// Keys should be custom types to avoid collisions in real code (string used here for simplicity).
func printValue(ctx context.Context) {
	if v := ctx.Value("userID"); v != nil {
		fmt.Printf("Context value for 'userID': %v\n", v)
	} else {
		fmt.Println("No value for 'userID' in context")
	}
}
