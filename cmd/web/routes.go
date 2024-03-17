package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// returns the mux
// makes the routes function a method of application
func (app *application) routes() http.Handler {

	// mux := http.NewServeMux()

	//pat is a third party library to create and handle routers/multiplexer
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/quote/create", http.HandlerFunc(app.createQuote))
	mux.Post("/quote/create", http.HandlerFunc(app.createQuote))//post request 
	mux.Get("/quote/:id", http.HandlerFunc(app.showQuote))
	// mux.HandleFunc("/show-quote", app.showQuotation)

	//create a file server to serve out static content
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))

	//create a variable to hold my middleware chain in order
	standardMiddleware := alice.New(
		app.recoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware,
	)

	return standardMiddleware.Then(mux)

	// return app.recoverPanicMiddleware(app.logRequestMiddleware(securityHeadersMiddleware(mux)))
}
