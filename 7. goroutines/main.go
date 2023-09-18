package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var mut sync.Mutex // pointer to a Mutex. mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex. A Mutex must not be copied after first use. it is used to synchronize access to a shared resource.
var wg sync.WaitGroup // pointer to a WaitGroup. A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

func main() {
	// go greeter("Hello")
	// greeter("Gophers")

	websiteList := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.stackoverflow.com/",
		"https://golang.org/",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1) // add 1 to the WaitGroup counter
	}

	wg.Wait() // blocks the execution of the program until the WaitGroup counter is 0
	fmt.Println(signals)
}


func getStatusCode(endpoint string) {
	defer wg.Done() // decrements the WaitGroup counter by 1

	res, err := http.Get(endpoint) // sends a GET request to the specified url

	if err != nil {
		fmt.Println(err)
		return
	}

	mut.Lock() // locks the Mutex. If the lock is already in use, the calling goroutine blocks until the mutex is available.
	signals = append(signals, endpoint)
	mut.Unlock() // unlocks the Mutex.

	fmt.Printf("Status code: %d for: %s\n", res.StatusCode, endpoint)
}