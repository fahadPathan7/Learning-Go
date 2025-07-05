package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
)

// 1. Basic Context Creation and Usage
func basicContextExample() {
    fmt.Println("=== Basic Context Examples ===")

    // Create a background context (root context)
    ctx := context.Background()
    fmt.Printf("Background context: %v\n", ctx)

    // Create a TODO context (placeholder when unsure what context to use)
    todoCtx := context.TODO()
    fmt.Printf("TODO context: %v\n", todoCtx)
}

// 2. Context with Timeout
func timeoutExample() {
    fmt.Println("\n=== Timeout Context Example ===")

    // Create context with 2-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Always call cancel to release resources

    // Simulate work that might take longer than timeout
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Work completed")
    case <-ctx.Done():
        fmt.Printf("Context cancelled: %v\n", ctx.Err())
    }
}

// 3. Context with Deadline
func deadlineExample() {
    fmt.Println("\n=== Deadline Context Example ===")

    // Create context with specific deadline
    deadline := time.Now().Add(1 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()

    // Check if context is done
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Work completed")
    case <-ctx.Done():
        fmt.Printf("Context deadline exceeded: %v\n", ctx.Err())
    }
}

// 4. Manual Cancellation
func cancellationExample() {
    fmt.Println("\n=== Manual Cancellation Example ===")

    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("Cancelling context...")
        cancel() // Manually cancel the context
    }()

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Work completed")
    case <-ctx.Done():
        fmt.Printf("Context manually cancelled: %v\n", ctx.Err())
    }
}

// 5. Context with Values
func valueExample() {
    fmt.Println("\n=== Context with Values Example ===")

    // Define custom key type to avoid key collisions
    type userIDKey string
    const userKey userIDKey = "userID"

    // Create context with value
    ctx := context.WithValue(context.Background(), userKey, "user123")

    // Retrieve value from context
    if userID := ctx.Value(userKey); userID != nil {
        fmt.Printf("User ID from context: %s\n", userID)
    }

    // Pass context to function
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    fmt.Println("Processing request with context...")
    type userIDKey string
    const userKey userIDKey = "userID"

    if userID := ctx.Value(userKey); userID != nil {
        fmt.Printf("Processing request for user: %s\n", userID)
    }
}

// 6. Context in HTTP Client
func httpClientExample() {
    fmt.Println("\n=== HTTP Client with Context Example ===")

    // Create context with timeout for HTTP request
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create HTTP request with context
    req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("HTTP request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("HTTP request completed with status: %s\n", resp.Status)
}

// 7. Context with Goroutines
func goroutineExample() {
    fmt.Println("\n=== Context with Goroutines Example ===")

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    var wg sync.WaitGroup

    // Start multiple workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d stopping: %v\n", id, ctx.Err())
            return
        case <-time.After(500 * time.Millisecond):
            fmt.Printf("Worker %d is working...\n", id)
        }
    }
}

// 8. Context Error Handling
func errorHandlingExample() {
    fmt.Println("\n=== Context Error Handling Example ===")

    // Timeout scenario
    ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel1()

    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Work completed")
    case <-ctx1.Done():
        err := ctx1.Err()
        switch err {
        case context.DeadlineExceeded:
            fmt.Println("Context deadline exceeded")
        case context.Canceled:
            fmt.Println("Context was cancelled")
        default:
            fmt.Printf("Unknown context error: %v\n", err)
        }
    }

    // Cancellation scenario
    ctx2, cancel2 := context.WithCancel(context.Background())
    cancel2() // Immediately cancel

    select {
    case <-ctx2.Done():
        if ctx2.Err() == context.Canceled {
            fmt.Println("Context was cancelled immediately")
        }
    default:
        fmt.Println("Context is still active")
    }
}

// 9. Context Best Practices Example
func bestPracticesExample() {
    fmt.Println("\n=== Context Best Practices Example ===")

    // Always pass context as first parameter
    processData(context.Background(), "sample data")

    // Chain contexts properly
    parentCtx, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer parentCancel()

    childCtx, childCancel := context.WithTimeout(parentCtx, 2*time.Second)
    defer childCancel()

    fmt.Printf("Parent deadline: %v\n", getDeadline(parentCtx))
    fmt.Printf("Child deadline: %v\n", getDeadline(childCtx))
}

func processData(ctx context.Context, data string) {
    // Always check context before expensive operations
    select {
    case <-ctx.Done():
        fmt.Printf("Processing cancelled: %v\n", ctx.Err())
        return
    default:
        fmt.Printf("Processing data: %s\n", data)
    }
}

func getDeadline(ctx context.Context) string {
    if deadline, ok := ctx.Deadline(); ok {
        return deadline.Format(time.RFC3339)
    }
    return "No deadline set"
}

// 10. Real-world Database Example
func databaseExample() {
    fmt.Println("\n=== Database Context Example ===")

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // Simulate database query with context
    result, err := queryDatabase(ctx, "SELECT * FROM users")
    if err != nil {
        log.Printf("Database query failed: %v", err)
        return
    }

    fmt.Printf("Database result: %s\n", result)
}

func queryDatabase(ctx context.Context, query string) (string, error) {
    // Simulate database operation
    select {
    case <-time.After(1 * time.Second):
        return "Query result", nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func main() {
    // basicContextExample()
    // timeoutExample()
    // deadlineExample()
    // cancellationExample()
    // valueExample()
    // httpClientExample()
    // goroutineExample()
    // errorHandlingExample()
    // bestPracticesExample()
    databaseExample()
}