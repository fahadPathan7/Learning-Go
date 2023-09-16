package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("learning json")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveHome")
	w.Write([]byte("<h1>hello world</h1>"))
}