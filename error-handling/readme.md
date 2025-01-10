# <div align='center'>Error Handling</div>

## Index
- [Error Handling](#error-handling)
  - [Index](#index)
  - [Introduction](#introduction)
  - [Some important error functions](#some-important-error-functions)
    - [1. `errors.New()`](#1-errorsnew)
    - [2. `fmt.Errorf()`](#2-fmterrorf)
    - [3. `errors.Is()`](#3-errorsis)
    - [4. `errors.As()`](#4-errorsas)
    - [5. `panic()`](#5-panic)
    - [6. `recover()`](#6-recover)
    - [7. `status.Errorf()`](#7-statuserrorf)
    - [8. `status.FromError()`](#8-statusfromerror)

---

## Introduction
Error handling is a mechanism to handle errors in a program. In Go, error handling is done using the `error` interface. It is a **built-in** interface that is defined in the `builtin` package. The `error` interface is defined as follows:

```go
type error interface {
    Error() string
}
```

The `error` interface has a single method `Error()`. This method returns a string that describes the error. If the method returns `nil`, then the error is considered to be nil.

In Go, errors are values. When a function returns an error, it is the responsibility of the caller to check the error and handle it appropriately. If the error is not handled, then it can lead to unexpected behavior in the program.

---

## Some important error functions

### 1. `errors.New()`
The `errors.New()` function is used to create a new error. It takes a string as an argument and returns a new error.

```go
package main

import (
    "errors"
    "fmt"
)

func errorFunc() error {
    return errors.New("This is an error")
}

func main() {
    err := errorFunc()
    if err != nil {
        fmt.Println(err) // Output: This is an error
    }
}
```

### 2. `fmt.Errorf()`
The `fmt.Errorf()` function is used to create a new error. It takes a format string and arguments as arguments and returns a new error.

```go
package main

import (
    "fmt"
)

func errorFunc() error {
    return fmt.Errorf("This is an error")
}

func main() {
    err := errorFunc()
    if err != nil {
        fmt.Println(err) // Output: This is an error
    }
}
```

### 3. `errors.Is()`
The `errors.Is()` function is used to check if an error is equal to another error. It takes two errors as arguments and returns a boolean value.

```go
package main

import (
    "errors"
    "fmt"
)

func errorFunc() error {
    return errors.New("This is an error")
}

func main() {
    err := errorFunc()
    if errors.Is(err, errors.New("This is an error")) {
        fmt.Println("Error is equal") // Output: Error is equal
    }
}
```

### 4. `errors.As()`
The `errors.As()` function is used to check if an error is of a specific type. It takes an error and a pointer to an interface as arguments and returns a boolean value.

```go
package main

import (
    "errors"
    "fmt"
)

func errorFunc() error {
    return errors.New("This is an error")
}

func main() {
    var err error
    errorFunc()
    if errors.As(err, &err) {
        fmt.Println("Error is of type error") // Output: Error is of type error
    }
}
```

### 5. `panic()`
The `panic()` function is used to stop the normal execution of a program. It takes an argument of any type and stops the execution of the program.

when a function calls `panic()`, the program stops executing the current function and starts unwinding the stack. It continues to unwind the stack until all the deferred functions are executed and then stops the program.

**Caution:** `panic()` should be used only in exceptional cases. It should not be used as a normal error handling mechanism.

```go
package main

import "fmt"

func riskyFunction() {
    fmt.Println("Before panic")
    panic("something went wrong") // This will stop the execution of the program
}
```

### 6. `recover()`
The `recover()` function is used to recover from a panic. It is used in combination with the `defer` statement.

```go
package main

import "fmt"

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Before panic")
    panic("something went wrong")
    fmt.Println("After panic") // This line will not be executed because of panic
}

func main() {
    fmt.Println("Before risky function")
    riskyFunction()
    fmt.Println("After risky function")
}
```

here, the `recover()` function is used to recover from the panic in the `riskyFunction()`. The `defer` statement is used to call the `recover()` function after the panic. The `recover()` function returns the value that was passed to the `panic()` function. If there was no panic, then the `recover()` function returns `nil`. The `recover()` function should be called inside a deferred function.

### 7. `status.Errorf()`
The `status.Errorf()` function is used to create a new error. It takes a `codes.Code` and a format string as arguments and returns a new error.

```go
package main

import (
    "fmt"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func errorFunc() error {
    return status.Errorf(codes.InvalidArgument, "This is an error")
}

func main() {
    err := errorFunc()
    if err != nil {
        fmt.Println(err) // Output: rpc error: code = InvalidArgument desc = This is an error
    }
}
```

benefits of using `status.Errorf()` over `fmt.Errorf()` is that it provides a standard way to create errors in gRPC services. It also provides a way to set the error code. The `codes.Code` is an enum that represents the error code.

### 8. `status.FromError()`
The `status.FromError()` function is used to get the status from an error. It takes an error as an argument and returns a `*status.Status`.

```go
package main

import (
    "fmt"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
)

func errorFunc() error {
    return status.Errorf(codes.InvalidArgument, "This is an error")
}

func main() {
    err := errorFunc()
    s, ok := status.FromError(err) // ok is true if the error is a gRPC status error
    if ok {
        fmt.Println(s.Code()) // Output: InvalidArgument
        fmt.Println(s.Message()) // Output: This is an error
    }
}
```

---