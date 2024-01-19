package main

import (
	//stuff specified here needs to be in alphabetical order
	"log"
	"net/http"
)

func main() {
	//create a file server(it loads the links to files if no .html is present)
	fileserver := http.FileServer(http.Dir(".")) //look in the current directory

	//create a web server responsible for listenting for requests
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", fileserver)
	log.Fatal(err)
}
