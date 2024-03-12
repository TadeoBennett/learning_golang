package main

import (
	"database/sql"
	"flag"
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
	//tell go to accepts the http adrress from the user
	// all commandline flags results will be stored in a pointerflag

	//create a commnadline flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

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

	//create a custom web server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(), //return the multiplexer
	}


	//if the flag is not provided it will use port :4000 by default as specified in the flag. 
	// To use another port, do the command: go run . -addr=":5000"
	//to let the user see how to use the flag, run the command:  go run . -addr
	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)

}
