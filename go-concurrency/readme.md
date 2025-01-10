# <div align='center'> Go Concurrency </div>

## Index
- [ Go Concurrency ](#-go-concurrency-)
  - [Index](#index)
  - [Go routines](#go-routines)
  - [Channels](#channels)
  - [Buffered Channels](#buffered-channels)
    - [Benefits of Buffered Channels](#benefits-of-buffered-channels)
    - [Size of Buffered Channels](#size-of-buffered-channels)
  - [Difference between buffered and unbuffered channels](#difference-between-buffered-and-unbuffered-channels)

<hr>

## Go routines
Go routines are lightweight threads of execution. They are used to run functions concurrently. Go routines are created using the `go` keyword followed by the function call.

Here is an example of creating a Go routine:
```go
import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go printNumbers()
    time.Sleep(5 * time.Second)
}
```
it will print numbers from 1 to 5 with a delay of 1 second between each number.

<hr>

## Channels
Channels are used to communicate between Go routines. They are used to **send and receive data between Go routines**. Channels are created using the `make` function.

Here is an example of channels:
```go
import (
    "fmt"
)

func sendNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i // send data to the channel
    }
    close(ch) // close the channel
}

func main() {
    ch := make(chan int) // create a channel
    go sendNumbers(ch)  // create a Go routine to send numbers
    for {
        num, ok := <-ch // receive data from the channel
        // check if the channel is closed
        if !ok {
            break
        }
        fmt.Println(num)
    }
}
```

In the above example, we have created a channel `ch` using the `make` function. We have created a Go routine `sendNumbers` to send numbers to the channel. In the `main` function, we are receiving data from the channel using the `<-` operator. We are also checking if the channel is closed using the `ok` variable.

<hr>

## Buffered Channels
Buffered channels are used to **store multiple values** in the channel. They have a buffer size that determines the number of values that can be stored in the channel.

Here is an example of buffered channels:
```go
import (
    "fmt"
)

func sendNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i // send data to the channel
    }
    close(ch) // close the channel
}

func main() {
    ch := make(chan int, 3) // create a buffered channel with a buffer size of 3
    go sendNumbers(ch)  // create a Go routine to send numbers
    for num := range ch {
        fmt.Println(num) // receive data from the channel
    }
}
```

In the above example, we have created a buffered channel `ch` with a buffer size of 3. We have created a Go routine `sendNumbers` to send numbers to the channel. In the `main` function, we are receiving data from the channel using the `range` keyword. The `range` keyword is used to iterate over the values in the channel.

### Benefits of Buffered Channels
- **Reduced Blocking**: Buffered channels can store multiple values, which reduces the chances of blocking the sender or receiver.
- **Improved Performance**: Buffered channels can improve the performance of the program by reducing the number of context switches between Go routines.
- **Asynchronous Communication**: Buffered channels allow asynchronous communication between Go routines, which can improve the overall efficiency of the program.
- **Reduced Latency**: Buffered channels can reduce the latency of the program by storing values in the buffer and processing them later.
- **Improved Scalability**: Buffered channels can improve the scalability of the program by allowing multiple Go routines to communicate efficiently.
- **Reduced Resource Consumption**: Buffered channels can reduce the resource consumption of the program by optimizing the communication between Go routines.

### Size of Buffered Channels
The size of the buffered channel determines the number of values that can be stored in the channel. If the buffer size is zero, the channel is unbuffered, and it can store only one value. If the buffer size is greater than zero, the channel is buffered, and it can store multiple values.

The buffer size should be chosen based on the requirements of the program. If the program requires high throughput and low latency, a larger buffer size can be used. If the program requires low latency and high concurrency, a smaller buffer size can be used.

The buffer size should be chosen based on the following factors:
- **throughput**: the amount of data that can be processed in a given time. It is a measure of the performance of the program. A higher throughput indicates better performance.
- **latency**: the time taken to process a request. It is a measure of the responsiveness of the program. A lower latency indicates better responsiveness.
- **concurrency**: the ability of the program to execute multiple tasks simultaneously. It is a measure of the efficiency of the program. Higher concurrency indicates better efficiency.

<hr>

## Difference between buffered and unbuffered channels

| Feature | Buffered Channels | Unbuffered Channels |
|---------|-------------------|---------------------|
| Blocking | Buffered channels do not block the sender until the buffer is full. | Unbuffered channels block the sender until the receiver is ready to receive the value. |
| Performance | Buffered channels can improve the performance of the program by reducing the number of context switches between Go routines. | Unbuffered channels can reduce the latency of the program by ensuring that the sender and receiver are synchronized. |
| Asynchronous Communication | Buffered channels allow asynchronous communication between Go routines. | Unbuffered channels ensure that the sender and receiver are synchronized. |
| Scalability | Buffered channels can improve the scalability of the program by allowing multiple Go routines to communicate efficiently. | Unbuffered channels can reduce the resource consumption of the program by optimizing the communication between Go routines. |
| Resource Consumption | Buffered channels can reduce the resource consumption of the program by storing values in the buffer and processing them later. | Unbuffered channels can reduce the resource consumption of the program by ensuring that the sender and receiver are synchronized. |
| Code | ```go ch := make(chan int, 3)``` | ```go ch := make(chan int)``` |

<hr>