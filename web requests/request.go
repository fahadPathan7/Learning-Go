package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.lco.dev"

func main() {
	fmt.Println("lco web requests")

	// by http.Get(url) we are sending a get request to the url.
	response, err := http.Get(url)

	// nil is used to represent the absence of a value
	if err != nil {
		panic(err) // panic is used to stop the execution of the program
	}

	fmt.Printf("Response is of type: %T\n", response) // response is of type: *http.Response. * is used to represent a pointer. So, response is a pointer to http.Response. This is because we are using the Get method of the http package. This method returns a pointer to http.Response. So, we are storing the pointer in response.

	defer response.Body.Close() // this means that we are done with the response body. So, we are closing it. Body is a stream of data. So, we are closing the stream of data. This is done to avoid memory leaks. defer is used to execute the statement at the end of the function. So, we are closing the response body at the end of the function.

	// reading the response body
	// we are using the ReadAll method of the ioutil package to read the response body
	// ReadAll method takes a reader as an argument and returns a byte slice and an error
	// response.Body is a reader
	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes) // converting the byte slice to a string
	fmt.Println(content)
}