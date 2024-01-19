package main

import (
	//stuff specified here needs to be in alphabetical order
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World. We are topot."))
}

func main() {
	//create a new serve mux
	mux := http.NewServeMux() //used to register all mapping between URLs and functions between mappings that will handle them

	//whenever the client types the root page, call the home function
	mux.HandleFunc("/", home)

	//create a web server responsible for listenting for requests
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
