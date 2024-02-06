package main

import (
	"html/template"
	"log"
	"net/http"
)

//The handler functions were moved here. You then just need to add "package main" 
//at the top of the file and the save the file and the dependencies gets added automatically

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Quotebox"))
}

func createQuoteForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/quotes_form_page.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//if there are no errors
	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
