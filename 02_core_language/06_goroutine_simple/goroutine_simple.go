// goroutine_simple.go
// Demonstrates basic usage of goroutines in Go.
//
// About goroutines:
// - Goroutines are lightweight, managed threads in Go.
// - Use 'go' keyword to start a new goroutine (concurrent function execution).
// - The Go runtime multiplexes goroutines onto OS threads.
// - Goroutines are much cheaper than OS threads (can run thousands).
//
// Comparison:
// - Python: Use threading.Thread for threads, or asyncio for async coroutines. Goroutines are more lightweight and easier to use than Python threads.
// - Java: Use Thread or ExecutorService for concurrency. Goroutines are more lightweight and managed by Go runtime, while Java threads are heavier and managed by JVM/OS.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Demonstrates goroutine basics, synchronization with sync.WaitGroup,
// a deliberate race condition (detectable with -race), and safe patterns
// using Mutex and channels.

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	fmt.Println("-- Goroutine basics and WaitGroup --")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg) // launch a goroutine with the 'go' keyword
	}
	wg.Wait()

	fmt.Println("\n-- Race condition example (run with -race to detect) --")
	// Unsafe increment: shared variable without synchronization -> data race
	counter := 0
	var wgRace sync.WaitGroup
	for i := 0; i < 100; i++ {
		wgRace.Add(1)
		go func() {
			defer wgRace.Done()
			counter++ // race: multiple goroutines write concurrently
		}()
	}
	wgRace.Wait()
	fmt.Println("counter (unsafely incremented) =", counter)

	fmt.Println("\n-- Fixing race with sync.Mutex --")
	var mu sync.Mutex
	safeCounter := 0
	var wgSafe sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wgSafe.Add(1)
		go func() {
			defer wgSafe.Done()
			mu.Lock()
			safeCounter++
			mu.Unlock()
		}()
	}
	wgSafe.Wait()
	fmt.Println("safeCounter (with Mutex) =", safeCounter)

	fmt.Println("\n-- Alternative: use a channel to serialize updates --")
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		sum := 0
		for v := range ch {
			sum += v
		}
		fmt.Println("sum via channel =", sum)
		close(done)
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	<-done

	fmt.Println()
	fmt.Println("Notes:")
	fmt.Println("- Goroutines are lightweight threads managed by Go's runtime.")
	fmt.Println("- The scheduler maps many goroutines onto OS threads efficiently.")
	fmt.Println("- Launch a goroutine by prefixing a call with 'go'.")
	fmt.Println("- Use sync.WaitGroup to wait for goroutines to finish.")
	fmt.Println("- Watch out for race conditions: use channels, Mutex, or atomic ops.")
	fmt.Println("- Detect races early with: go run -race goroutine_simple.go")
}
