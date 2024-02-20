package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"tadeobennett.net/quotation/pkg/models/postgresql"

)

// provide the credentials for our database
const (
	host     = "localhost"
	port     = 5432
	user     = "quotebox"
	password = "tadeo2002"
	dbname   = "quotebox"
)

func setUpDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err //return nil and the error
	}

	err = db.Ping() //Test our connection
	if err != nil {
		return nil, err
	}
	fmt.Println("Database Connection established")
	return db, nil //return the connection and nil(no errors)
}

// Dependencies (things/variables)
// DEPENDENCY INJECTION
type application struct {
	quotes *postgresql.QuoteModel //references the QuoteModel which has the db connection
}

func main() {
	//FIRST CONNECT TO THE DATABASE --------------------------------
	var db, err = setUpDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() //can only close the connection where the function is called

	//passing the database connection to a model so it can handle other database operations
	app := &application{
		quotes: &postgresql.QuoteModel{
			DB: db,
		},
	}

	//SECOND CREATE THE SERVER INSTANCE ----------------------------------
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/quote", app.createQuoteForm)
	mux.HandleFunc("/quote-add", app.createQuote)
	mux.HandleFunc("/show", app.displayQuotation)
	log.Println("Starting a server on port :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
