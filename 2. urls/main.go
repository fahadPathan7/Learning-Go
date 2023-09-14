package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=reactjs&paymentid=1234&price=50"

func main() {
	fmt.Println("handling urls")
	fmt.Println(myurl)

	// parse url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme) // https. scheme is protocol
	fmt.Println(result.Host)   // lco.dev:3000. host is domain name and port
	fmt.Println(result.Port()) // 3000.
	fmt.Println(result.Path)   // /learn.
	fmt.Println(result.RawQuery) // course=reactjs&paymentid=1234&price=50


	// query params
	queryparams := result.Query()
	fmt.Printf("type of queryparams is %T\n", queryparams) //url.Values. key value pair

	fmt.Println(queryparams["course"]) // [reactjs]

	// looping over query params
	for key, value := range queryparams {
		fmt.Println("Param ", key, ": ", value)
	}

	// creating a new url
	newurl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	fmt.Println(newurl.String()) // https://lco.dev/tutcss?user=hitesh
}