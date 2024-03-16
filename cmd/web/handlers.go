package main

import (
	// "fmt"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

//moved the Quotation struct to another location for better organization

//The handler functions were moved here. You then just need to add "package main"
//at the top of the file and the save the file and the dependencies gets added automatically

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) //use our custom log
		return
	}
	w.Write([]byte("Welcome to Quotebox"))

}

// these are now hanlder methods of the application type and not handler functions after adding "(app *application)"
func (app *application) createQuoteForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../ui/html/quotes_form_page.tmpl")

	if err != nil {
		log.Println(err.Error())
		app.serverError(w, err)
		return
	}

	//if there are no errors
	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) createQuote(w http.ResponseWriter, r *http.Request) {
	//go back to the form if this location is not accessed through the post method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/quote", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//get the values from the request
	author := r.PostForm.Get("author_name")
	category := r.PostForm.Get("category")
	quote := r.PostForm.Get("quote")

	//check the web form fields for validity. We will use a map to save the errrors
	// errors := make(map[typeofKEY]typeofVALUE)
	errors := make(map[string]string)

	//check each field
	if strings.TrimSpace(author) == "" {
		errors["author"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(author) > 10 { //RunCountInString is used to count the characters
		errors["author"] = "This field is too long (max 50 characters)"
	}

	if strings.TrimSpace(category) == "" {
		errors["category"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(category) > 10 { //RunCountInString is used to count the characters
		errors["category"] = "This field is too long (max 50 characters)"
	}

	if strings.TrimSpace(quote) == "" {
		errors["quote"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(quote) > 20 { //RunCountInString is used to count the characters
		errors["quote"] = "This field is too long (max 50 characters)"
	}

	if len(errors) > 0 { //an error exists
		ts, err := template.ParseFiles("../../ui/html/quotes_form_page.tmpl") //load the template file

		if err != nil { //error loading the template
			log.Println(err.Error())
			app.serverError(w, err)
			return
		}

		err = ts.Execute(w, &templateData{
			ErrorsFromForm: errors,
			FormData:       r.PostForm,
		})
		if err != nil {
			log.Panicln(err.Error())
			app.serverError(w, err)
			return
		}

		return
	}

	//insert a quote
	id, err := app.quotes.Insert(author, category, quote)

	//check if an error was returned from the insert function
	if err != nil {
		log.Println(err.Error())
		app.serverError(w, err)
		return
	}

	//won't show on the page because the page has been redirected
	fmt.Fprintf(w, "row with id %d has been inserted.", id)

	http.Redirect(w, r, "/show", http.StatusSeeOther)

}

func (app *application) displayQuotation(w http.ResponseWriter, r *http.Request) {
	q, err := app.quotes.Read()

	if err != nil {
		log.Println(err.Error())
		app.serverError(w, err)
		return
	}

	//an instance of template data -------------------------------
	data := &templateData{
		Quotes: q,
	}

	//Display quotes using a template
	ts, err := template.ParseFiles("../../ui/html/show_page.tmpl")

	if err != nil {
		log.Println(err.Error())
		app.serverError(w, err)
		return
	}

	//if there are no errors
	err = ts.Execute(w, data)

	if err != nil {
		log.Panicln(err.Error())
		app.serverError(w, err)
	}
}
