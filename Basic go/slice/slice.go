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

	slice := []int{1, 2, 3, 4, 5}
	modifySlice(slice)
	fmt.Println(slice)
}

func modifySlice(slice []int) {
	slice[0] = 100
}