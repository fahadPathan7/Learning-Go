// mongodb+srv://fahadpathan56:<password>@cluster0.oxoqi6z.mongodb.net/?retryWrites=true&w=majority

package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)


func main() {
	fmt.Println("server is starting at port 5000")
	r := router.Router() // create router.
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("server is running at port 5000")
}