# Concurrency

[← Back to Index](README.md) | [Next: Modern Tooling →](10_modern_tooling.md)

Go's concurrency model requires discipline.

## Channel Axioms and Patterns

### 1. Channels Are for Communication

```go
// Bad: Using channels for synchronization
func worker(done chan bool) {
    // do work
    done <- true  // Just for signaling
}

// Better: Use sync.WaitGroup for synchronization
func worker(wg *sync.WaitGroup) {
    defer wg.Done()
    // do work
}
```

### 2. Buffered Channel Sizing

```go
// Unbuffered: Synchronous communication
ch := make(chan int)  // Size 0

// Size 1: Non-blocking send for common case
ch := make(chan int, 1)

// Avoid large buffers: they hide backpressure issues
ch := make(chan int, 1000)  // Why 1000? Usually a smell
```

### 3. Keep Busy or Do the Work

```go
// Bad: Spawning goroutine for trivial work
func process(data []byte) {
    go func() {
        result := expensiveWork(data)
        // now what? Hard to get result back
    }()
}

// Good: Let caller decide concurrency
func process(data []byte) Result {
    return expensiveWork(data)
}

// Caller can use concurrency if needed
go func() {
    result := process(data)
    // handle result
}()
```

### 4. Always Know When Goroutines Stop

```go
// Bad: Goroutine leak
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever if nobody sends
        process(val)
    }()
    // Goroutine never exits!
}

// Good: Use context for cancellation
func noLeak(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            process(val)
        case <-ctx.Done():
            return  // Goroutine can exit
        }
    }()
}
```

### 5. Context for Cancellation (Standard Pattern)

```go
import "context"

func processWithTimeout(ctx context.Context, data []byte) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    resultCh := make(chan error, 1)
    go func() {
        resultCh <- doWork(data)
    }()
    
    select {
    case err := <-resultCh:
        return err
    case <-ctx.Done():
        return ctx.Err()  // timeout or cancelled
    }
}
```

### 6. errgroup for Goroutine Management

```go
import "golang.org/x/sync/errgroup"

func processBatch(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)
    
    for _, item := range items {
        item := item  // Capture loop variable
        g.Go(func() error {
            return processItem(ctx, item)
        })
    }
    
    return g.Wait()  // Waits for all goroutines, returns first error
}
```

### 7. Worker Pool Pattern

```go
func workerPool(ctx context.Context, jobs <-chan Job) error {
    const numWorkers = 5
    g, ctx := errgroup.WithContext(ctx)
    
    for i := 0; i < numWorkers; i++ {
        g.Go(func() error {
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return nil  // Channel closed
                    }
                    if err := job.Process(); err != nil {
                        return err
                    }
                case <-ctx.Done():
                    return ctx.Err()
                }
            }
        })
    }
    
    return g.Wait()
}
```

### 8. Type-Safe Channels with Generics (Go 1.18+)

```go
// Generic pipeline stage
func pipeline[T, U any](ctx context.Context, in <-chan T, fn func(T) U) <-chan U {
    out := make(chan U)
    go func() {
        defer close(out)
        for {
            select {
            case v, ok := <-in:
                if !ok {
                    return
                }
                select {
                case out <- fn(v):
                case <-ctx.Done():
                    return
                }
            case <-ctx.Done():
                return
            }
        }
    }()
    return out
}
```

## Concurrency Checklist

- ✅ Use channels for communication, not synchronization
- ✅ Keep channel buffers small (0 or 1 typically)
- ✅ Let callers control concurrency
- ✅ Always know when goroutines will exit
- ✅ Use context for cancellation and timeouts
- ✅ Use errgroup for managing multiple goroutines
- ✅ Avoid goroutine leaks with proper cleanup
- ✅ Test concurrent code with `-race` flag

## Common Pitfalls

**Goroutine Leaks:**
- Forgotten select with `<-ctx.Done()`
- Channels that never close
- Waiting on channels that no one writes to

**Race Conditions:**
- Shared state without proper synchronization
- Capturing loop variables incorrectly
- Reading/writing without locks or atomic operations

**Deadlocks:**
- Circular waiting on channels
- Forgetting to send/receive on channels
- Improper lock ordering

---

[← Previous: Testing](08_testing.md) | [Next: Modern Tooling →](10_modern_tooling.md)
