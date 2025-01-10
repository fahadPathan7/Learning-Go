package main

import "fmt"

func riskyFunction2() error {
	return fmt.Errorf("something went wrong")
}

func main2() {
	err := riskyFunction2()
	if err != nil {
		fmt.Println("Error:", err)
	}
}