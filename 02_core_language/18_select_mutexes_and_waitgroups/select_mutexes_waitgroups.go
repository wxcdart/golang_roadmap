package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Demonstrates using select to handle multiple channels, sync.WaitGroup to
// wait for goroutines, sync.Mutex to protect shared state, and a channel-based
// alternative for serialized updates. Also describes when to prefer channels vs locks.

func worker(id int, results chan<- int, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	// simulate variable work
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	if rand.Intn(8) == 0 { // occasional error
		errs <- fmt.Errorf("worker %d failed", id)
		return
	}
	results <- id * 10
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Channels used to collect results and errors
	results := make(chan int)
	errs := make(chan error)

	var wg sync.WaitGroup
	workerCount := 6
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, results, errs, &wg)
	}

	// close a 'done' channel when all workers finish
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Use select to receive from multiple channels and handle timeouts
	fmt.Println("-- receiving results and errors via select --")
	for {
		select {
		case r := <-results:
			fmt.Println("result:", r)
		case e := <-errs:
			fmt.Println("error:", e)
		case <-done:
			fmt.Println("all workers completed")
			// drain any remaining results/errors non-blockingly
			drain := true
			for drain {
				select {
				case r := <-results:
					fmt.Println("(drain) result:", r)
				case e := <-errs:
					fmt.Println("(drain) error:", e)
				default:
					drain = false
				}
			}
			goto afterSelect
		case <-time.After(300 * time.Millisecond):
			fmt.Println("waiting... (no activity for 300ms)")
		}
	}

afterSelect:
	// --- Mutex example: protecting shared map state ---
	fmt.Println("\n-- Mutex protecting shared state example --")
	var mu sync.Mutex
	counts := make(map[int]int)
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				counts[id]++
				mu.Unlock()
			}
		}(i)
	}
	wg2.Wait()
	total := 0
	for k, v := range counts {
		fmt.Printf("counts[%d]=%d\n", k, v)
		total += v
	}
	fmt.Println("total (mutex-protected) =", total)

	// --- Channel-serialized updates (alternative to locks) ---
	fmt.Println("\n-- Channel-serialized updates (no locks) --")
	updates := make(chan int)
	doneSerial := make(chan struct{})
	go func() {
		sum := 0
		for v := range updates {
			sum += v
		}
		fmt.Println("sum via updates channel =", sum)
		close(doneSerial)
	}()

	// multiple goroutines send increments into the updates channel
	var wg3 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for k := 0; k < 10; k++ {
				updates <- 1
			}
		}()
	}
	wg3.Wait()
	close(updates)
	<-doneSerial

	// Guidance
	// Decision table (from attached image):
	// | Situation                                           | Use       |
	// |-----------------------------------------------------|-----------|
	// | You're passing data between goroutines              | Channels  |
	// | You're protecting shared state (e.g., a counter)    | Mutex     |
	// | You need to wait for tasks to complete              | WaitGroup |
	// | You're monitoring multiple operations               | select    |
	//
	fmt.Println("\nGuidance: when to use channels vs Mutex:")
	fmt.Println("- Use channels when you're passing data between goroutines or want serialized handling of events.")
	fmt.Println("- Use Mutex when protecting complex shared state (maps, multi-field structs) where channel-based design would be awkward or inefficient.")
	fmt.Println("- Channels make coordination easier; Mutexes are lower overhead for simple shared-state updates.")
	fmt.Println("- The attachment summary is useful: Channels = passing data, Mutex = protecting shared state, WaitGroup = wait for tasks, select = monitor multiple channels.")
}
