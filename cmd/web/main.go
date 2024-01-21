// Program Demonstrating:
// 1. Form Building
// 2. Redirecting
// 3. Providing multiple pieces of data to HTML templates 
// TASK: Calculate and display the area of a rectangle


package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World. We are topot."))
}

func displayTime(w http.ResponseWriter, r *http.Request) {
	localTime := time.Now().Format("3:04:05 PM")
	ts, _ := template.ParseFiles("./ui/html/display.time.page.tmpl")
	ts.Execute(w, localTime)
}

//show the form to the user
func getValues(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, _ := template.ParseFiles("./ui/html/input.area.page.tmpl")
		ts.Execute(w, nil)
	}else{
		http.Redirect(w, r, "/area-calculator-2", http.StatusTemporaryRedirect)
	}
}

//compute and display the area
func calculateArea(w http.ResponseWriter, r *http.Request) {
	//sending multiple pieces of data using a struct
	type UserData struct{
		Length float64
		Width float64
		Area float64
	}

	//get the length and width from the form
	r.ParseForm()

	//save the values
	length := r.PostForm.Get("length")
	width := r.PostForm.Get("width")

	//calculate the area (also converting the values to number types[musit import strconv])
	lengthOfRectangle, _ := strconv.ParseFloat(length, 64)
	widthOfRectangle, _ := strconv.ParseFloat(width, 64)
	areaOfRectangle := lengthOfRectangle * widthOfRectangle

	//create and instance of the UserData
	data := UserData {
		lengthOfRectangle,
		widthOfRectangle,
		areaOfRectangle,
	}

	//call the template engine
	ts, _ := template.ParseFiles("./ui/html/display.area.page.tmpl")
	ts.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/time", displayTime)
	mux.HandleFunc("/area-calculator", getValues)
	mux.HandleFunc("/area-calculator-2", calculateArea)
	fileServer := http.FileServer(http.Dir("./ui/static/")) //strip prefix from url to get the right name to send to the file server because the url is appended to the directory location provided to the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
