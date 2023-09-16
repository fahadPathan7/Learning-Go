package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func performGetRequest() {
	fmt.Println("\nPerforming GET Request")

	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		fmt.Println("Error performing GET request:", err)
		panic(err)
	}

	defer response.Body.Close()

	// Print the response status, body and headers
	fmt.Println("Response status:", response.Status)
	fmt.Println("Response body:", response.Body)
	fmt.Println("Response headers:", response.Header)

}

func performPostRequest() {
	fmt.Println("\nPerforming POST Request")

	const myUrl = "http://localhost:8000/post"

	// Create a new POST request with JSON data. The request body is of type io.Reader. We can use strings.NewReader to initialize it from a string.
	requestBody := strings.NewReader(`
		{
			"coursename": "golang",
			"price": 0,
			"platform": "youtube"
		}
	`)

	// Perform the request. http.Post is a shortcut for http.Client.Post. application/json is the most common POST request Content-Type. The third argument is the request body.
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		fmt.Println("Error performing POST request:", err)
		panic(err)
	}

	defer response.Body.Close()

	// ioutil.ReadAll reads the response body into a byte slice. We can then convert it to a string.
	content, _ := ioutil.ReadAll(response.Body)

	// printing after converting byte slice to string
	fmt.Println(string(content))
}

func main() {
	performGetRequest()
	performPostRequest()
}