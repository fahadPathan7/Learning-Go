package main

import (
    "fmt"
)

type MyError struct {
    message string
}

func (e *MyError) Error() string {
    return e.message
}

func errorFunc() error {
    return &MyError{"This is an error"}
}

func main() {
    err := errorFunc()
    if err != nil {
        fmt.Println(err) // Output: This is an error
    }
}