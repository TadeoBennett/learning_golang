package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Quotation struct {
	Quotation_id   int
	Insertion_date time.Time
	Author_name    string
	Category       string
	Quote          string
}

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

func (app *application) displayQuotation(w http.ResponseWriter, r *http.Request) {

	readQuotes := `
	SELECT *
	FROM quotations
	LIMIT 5
	`

	rows, err := app.db.Query(readQuotes) //returns the rows of results
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	defer rows.Close()

	//Store the quotations in a slice of Quotation (struct)
	var quotes []Quotation

	//Iterate over rows (a result set)
	for rows.Next() {
		//Create a Quotation for the row
		var q Quotation

		err = rows.Scan(&q.Quotation_id, &q.Insertion_date, &q.Author_name, &q.Category, &q.Quote)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		//Append to quotes
		quotes = append(quotes, q)
	}

	//Always check the rows.Err()
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	//Print our quotes
	for _, quote := range quotes {
		// fmt.Printf("ID: %d, Date: %s, Author: %s, Category: %s, Quote: %s\n",
		// 	quote.Quotation_id, quote.Insertion_date, quote.Author_name, quote.Category, quote.Quote)
		fmt.Fprintf(w, "%v\n", quote)
	}
}
