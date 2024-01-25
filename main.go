package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //third party package
	"log"
)

// provide the credentials for our database
const (
	host     = "localhost"
	port     = 5432
	user     = "quotebox"
	password = "tadeo2002"
	dbname   = "quotebox"
)

// dsn: data source name
func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() //needs to be closed to free space

	err = db.Ping() //Test our connection
	if err != nil {
		log.Fatal(err)
	}

	//Insert a quote
	insertQuote := `
	INSERT INTO quotations(author_name, category, quote)
	VALUES
	($1, $2, $3);
	`

	//note that Exec() does not return anything
	_, err = db.Exec(insertQuote,
		"Lao Tzu",
		"Life",
		"Nature does not hurry, yet everything is accomplished.")

	//if the query did not work
	if err != nil {
		log.Fatal(err)
	}

	//query to the primary key of a recently inserted value (The returning command)
	// INSERT INTO quotations (author_name, category, quote)
	// VALUES
	// ('Lao Tzu', 'Life', 'The best fighter is never angry')
	// RETURNING quotation_id;

	//query to get the insertion date of a recent value generated by postgres
	insertQuote2 := `
	INSERT INTO quotations(author_name, category, quote)
	VALUES
	($1, $2, $3)
	RETURNING quotation_id;	
	`

	quotation_id := 0; //every variable returned needs a value to store it
	//note that QueryRow returns info about the values
	err = db.QueryRow(insertQuote2,
		"Lao Tzu",
		"Life",
		"Mastering others is strength. Mastering yourself is true power.").Scan(&quotation_id)

	//if the query did not work
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The recently inserted record has quotation_id: ", quotation_id)

}
