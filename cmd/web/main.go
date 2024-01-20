package main

import (
	//stuff specified here needs to be in alphabetical order

	"html/template"
	"log"
	"net/http"
	"time"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World. We are topot."))
}

func displayTime(w http.ResponseWriter, r *http.Request) {
	//step1: get the time
	localTime := time.Now().Format("3:04:05 PM")

	//step2: read in the template file
	ts, _ := template.ParseFiles("./ui/html/display.time.tmpl")

	//step3: do the substitution (template engine)
	ts.Execute(w, localTime)
}

func main() {
	//create a new serve mux
	mux := http.NewServeMux() //used to register all mapping between URLs and functions between mappings that will handle them

	//[for handling FUNCTIONS]whenever the client types the root page, call the home function
	mux.HandleFunc("/", home)

	//The URL for the time
	mux.HandleFunc("/time", displayTime)

	//create a file server
	fileServer := http.FileServer(http.Dir("./ui/static/")) //strip prefix from url to get the right name to send to the file server because the url is appended to the directory location provided to the file server

	//create a URL mapping to the static directory
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	//create a web server responsible for listenting for requests
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
