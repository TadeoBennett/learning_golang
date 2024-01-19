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

	//[for handling FUNCTIONS]whenever the client types the root page, call the home function
	mux.HandleFunc("/", home)

	//create a file server
	fileServer := http.FileServer(http.Dir("./static"))//strip prefix from url to get the right name to send to the file server because the url is appended to the directory location provided to the file server

	//create a URL mapping to the static directory
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	//create a web server responsible for listenting for requests
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
