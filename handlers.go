package main

import (
	"html/template"
	"log"
	"net/http"
)

//The handler functions were moved here. You then just need to add "package main"
//at the top of the file and the save the file and the dependencies gets added automatically

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Quotebox"))
}

// these are now hanlder methods of the application type and not handler functions after adding "(app *application)"
func (app *application) createQuoteForm(w http.ResponseWriter, r *http.Request) {
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

func (app *application) createQuote(w http.ResponseWriter, r *http.Request) {
	//go back to the form if this location is not accessed through the post method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/quote", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	author := r.PostForm.Get("author_name")
	category := r.PostForm.Get("category")
	quote := r.PostForm.Get("quote")

	s := `
	INSERT INTO quotations(author_name, category, quote)
	VALUES ($1, $2, $3)
	`

	_, err = app.db.Exec(s, author, category, quote)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
