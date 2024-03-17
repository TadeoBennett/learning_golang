package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// returns the mux
// makes the routes function a method of application
func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/quote", app.createQuoteForm)
	mux.HandleFunc("/quote-add", app.createQuote)
	mux.HandleFunc("/show", app.displayQuotation)

	//create a file server to serve out static content
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	//create a variable to hold my middleware chain in order
	standardMiddleware := alice.New(
		app.recoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware,
	)

	return standardMiddleware.Then(mux)

	// return app.recoverPanicMiddleware(app.logRequestMiddleware(securityHeadersMiddleware(mux)))
}
