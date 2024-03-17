package main

import (
	// "fmt"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"tadeobennett.net/quotation/pkg/models"
)

// displays all the quotes
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w) //use our custom log
	// 	return
	// }
	//the above is taken care of in the routes.go using pat (it only looks for "/")

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

func (app *application) createQuoteForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../ui/html/quotes_form_page.tmpl")

	if err != nil {
		app.serverError(w, err)
		return
	}

	//if there are no errors
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createQuote(w http.ResponseWriter, r *http.Request) {
	//go back to the form if this location is not accessed through the post method
	// if r.Method != http.MethodPost {
	// 	http.Redirect(w, r, "/quote", http.StatusSeeOther)
	// 	return
	// }
	//the above is only called when we are using post so the code is not necessary

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
	} else if utf8.RuneCountInString(author) > 25 { //RunCountInString is used to count the characters
		errors["author"] = "This field is too long (max 25 characters)"
	}

	if strings.TrimSpace(category) == "" {
		errors["category"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(category) > 15 { //RunCountInString is used to count the characters
		errors["category"] = "This field is too long (max 15 characters)"
	}

	if strings.TrimSpace(quote) == "" {
		errors["quote"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(quote) > 50 { //RunCountInString is used to count the characters
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

	http.Redirect(w, r, fmt.Sprintf("/quote/%d", id), http.StatusSeeOther)

}

func (app *application) showQuote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((r.URL.Query().Get(":id"))) //gets the ID from the URL

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	q, err := app.quotes.Get(id)

	//check the errors returned
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
	}

	fmt.Fprintf(w, "%v", q)
}
