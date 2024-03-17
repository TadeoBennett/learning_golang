package main

import (
	"net/http"
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

	//the initial request will be passed to the log
	return app.recoverPanicMiddleware(app.logRequestMiddleware(securityHeadersMiddleware(mux))) 
}
