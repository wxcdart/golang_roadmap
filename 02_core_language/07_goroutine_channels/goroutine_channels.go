// goroutine_channels.go
// Demonstrates basic usage of goroutines and channels in Go.
// Channels are the primary way to safely communicate between goroutines.
// This example includes comments and cross-language comparisons (Python/Java).

package main

import (
	"fmt"
	"sync"
	"time"
)

// Demonstrates channel basics, buffered vs unbuffered behavior,
// sending/receiving, closing channels, and channel direction types.

// receiveOnly demonstrates a function that only receives from a channel.
func receiveOnly(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch { // ranges until channel is closed
		fmt.Printf("receiver %d got %d\n", id, v)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Printf("receiver %d done (channel closed)\n", id)
}

// sendOnly demonstrates a function that only sends to a channel.
func sendOnly(ch chan<- int, nums []int) {
	for _, n := range nums {
		ch <- n // send will block if channel is unbuffered and no receiver ready
		fmt.Printf("sent %d\n", n)
	}
	// Only the sender should close the channel when done
	close(ch)
}

func main() {
	fmt.Println("-- Unbuffered channel (synchronous) --")
	unbuf := make(chan int) // unbuffered
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOnly(1, unbuf, &wg)

	// Send and receive must happen at the same time for unbuffered channels;
	// the send will block until the receiver is ready.
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("sending %d to unbuffered channel\n", i)
			unbuf <- i
			fmt.Printf("sent %d to unbuffered channel\n", i)
		}
		close(unbuf)
	}()
	wg.Wait()

	fmt.Println("\n-- Buffered channel (asynchronous up to capacity) --")
	buf := make(chan int, 2) // buffered with capacity 2
	wg.Add(1)
	go receiveOnly(2, buf, &wg)

	// These sends will only block if buffer is full
	go func() {
		for i := 10; i <= 15; i++ {
			fmt.Printf("attempting send %d to buffered channel\n", i)
			buf <- i
			fmt.Printf("sent %d to buffered channel\n", i)
		}
		close(buf)
	}()
	wg.Wait()

	fmt.Println("\n-- Using send-only and receive-only channel direction in APIs --")
	nums := []int{100, 200, 300}
	ch := make(chan int)
	wg.Add(1)
	go receiveOnly(3, ch, &wg)
	sendOnly(ch, nums) // sendOnly closes the channel
	wg.Wait()

	fmt.Println("\nNotes:")
	fmt.Println("- Channels connect goroutines and synchronize execution.")
	fmt.Println("- Unbuffered channels block the sender until a receiver is ready.")
	fmt.Println("- Buffered channels allow limited asynchronous sends up to their capacity.")
	fmt.Println("- Close channels from the sender side to signal no more values.")
	fmt.Println("- Use channel directions (chan<-, <-chan) to enforce safe APIs.")
}
