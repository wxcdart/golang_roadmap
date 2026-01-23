// concurrent_workers.go
// Demonstrates a worker pool using goroutines, channels, context, and sync.WaitGroup.
//
// This example shows:
// - Creating a worker pool with goroutines
// - Sending jobs and collecting results via channels
// - Using context for cancellation and timeouts
// - Waiting for all workers to finish with sync.WaitGroup

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a unit of work
// -----------------------------
// NOTE: In Go, struct types must be declared with 'type' and curly braces.
// The previous version was missing 'type' and the braces, causing a syntax error.
// Example of correct struct declaration:
//
//	type Job struct {
//	    ID int // Unique job identifier
//	}
type Job struct {
	ID int // Unique job identifier
}

// Result represents the result of a job
// -------------------------------------
// The same applies for the Result struct. Always use 'type' and curly braces.
type Result struct {
	JobID int // ID of the job processed
	Value int // Result value (e.g., job.ID * 2)
}

// worker processes jobs from the jobs channel and sends results to the results channel.
// It listens for context cancellation and exits if the context is done.
func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this worker as done when function exits
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelled\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			// Simulate work by sleeping for a random duration
			duration := time.Duration(rand.Intn(500)+100) * time.Millisecond
			time.Sleep(duration)
			// Send result to results channel
			results <- Result{JobID: job.ID, Value: job.ID * 2}
			fmt.Printf("Worker %d: processed job %d\n", id, job.ID)
		}
	}
}

func main() {
	numWorkers := 3 // Number of worker goroutines
	numJobs := 10   // Number of jobs to process

	jobs := make(chan Job)       // Channel for sending jobs to workers
	results := make(chan Result) // Channel for collecting results from workers
	var wg sync.WaitGroup        // WaitGroup to wait for all workers to finish

	// Create a context with timeout for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	// Send jobs to the jobs channel in a separate goroutine
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j}
		}
		close(jobs) // Close jobs channel when done
	}()

	// Collect results in a separate goroutine
	done := make(chan struct{})
	go func() {
		for r := 1; r <= numJobs; r++ {
			select {
			case res := <-results:
				fmt.Printf("Result: job %d -> %d\n", res.JobID, res.Value)
			case <-ctx.Done():
				fmt.Println("Result collection cancelled")
				return
			}
		}
		done <- struct{}{} // Signal that all results are collected
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close results channel after all workers are done

	// Wait for results to be collected or context to timeout
	select {
	case <-done:
		fmt.Println("All results collected.")
	case <-ctx.Done():
		fmt.Println("Main: context timeout or cancelled.")
	}
}
