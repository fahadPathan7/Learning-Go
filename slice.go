package main

import "fmt"

func main() {
	// this is a dynamic array
	var names []string

	var studentCnt int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&studentCnt)

	for i := 0; i < studentCnt; i++ {
		var name string
		fmt.Print("Enter name: ")
		fmt.Scan(&name)
		names = append(names, name)
	}

	fmt.Println("Number of students: ", len(names))
	fmt.Println(names)
}