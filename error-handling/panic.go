package main

import "fmt"

func riskyFunction1() {
	// Defer is used to execute a function after the surrounding function returns
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("something went wrong")
	fmt.Println("After panic") // This line will not be executed because of panic
}

func main1() {
	fmt.Println("Before risky function")
	riskyFunction1()
	fmt.Println("After risky function")
}
