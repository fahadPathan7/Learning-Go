package main

import (
	"fmt"
	"time"
	//"sync"
)

func sayHello() {
	fmt.Println("Hello")
}

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	time.Sleep(2 * time.Second)
	fmt.Println("Bye")
}

func sum(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	ch := make(chan int)
	go sum(1, 2, ch)
	go sum(3, 4, ch)
	x, y := <-ch, <-ch // receive from c (order is not guaranteed)
	fmt.Println(x, y)
	go sayHello()
	go sayHi()
	go sayBye()
	time.Sleep(5 * time.Second)
}