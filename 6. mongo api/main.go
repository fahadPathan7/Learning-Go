// mongodb+srv://fahadpathan56:<password>@cluster0.oxoqi6z.mongodb.net/?retryWrites=true&w=majority

package main

import (
	"fmt"
	//"log"
	"mongodb/router"
	"net/http"

	"github.com/rs/cors"
)


func main() {
	fmt.Println("server is starting at port 5000")
	r := router.Router() // create router.
	// Create a CORS handler with desired options
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Replace with your allowed origins
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })

    // Wrap your router with the CORS handler
    handler := c.Handler(r)

    http.Handle("/", handler)
    http.ListenAndServe(":5000", nil)
	fmt.Println("server is running at port 5000")
}