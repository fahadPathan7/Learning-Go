package main

import (
	"errors"
	"fmt"
)

func riskyFunction3() error {
	errorMessage := errors.New("something went wrong")
	return errorMessage
}

func main3() {
	err := riskyFunction3()
	if err != nil {
		fmt.Println("Error:", err)
	}
}