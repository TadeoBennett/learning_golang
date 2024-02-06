package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// provide the credentials for our database
const (
	host     = "localhost"
	port     = 5432
	user     = "quotebox"
	password = "tadeo2002"
	dbname   = "quotebox"
)

func setUpDB()(*sql.DB, error){
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


//this is a global variable approach to connecting to the db to perform crud operations
//the shorthand notation := does not work outside a function
//since db needs to be used in two functions we write it outside
var db, err = setUpDB()

func main() {
	//FIRST CONNECT TO THE DATABASE --------------------------------
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() //can only close the connection where the function is called

	//SECOND CREATE THE SERVER INSTANCE ----------------------------------
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/quote", createQuoteForm)
	mux.HandleFunc("/quote-add", createQuote)
	log.Println("Starting a server on port :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
