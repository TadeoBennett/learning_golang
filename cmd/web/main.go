package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"tadeobennett.net/quotation/pkg/models/postgresql"
)

func setUpDB(dsn string) (*sql.DB, error) {
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

func loadEnvVariables() {
	envFilePath := "../../.env"

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type application struct {
	quotes *postgresql.QuoteModel //references the QuoteModel which has the db connection
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//tell go to accepts the http adrress from the user
	// all commandline flags results will be stored in a pointerflag

	//create a commandline flag
	//to edit the provided address using the commandline, use 'go run . -addr=":number"'
	addr := flag.String("addr", ":4000", "HTTP network address")

	loadEnvVariables() //loads the variables in the .env file

	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")

	//now you can provide a new dsn flag for the connection
	dsn := flag.String("dsn",
		"postgres://"+dbname+":"+password+"@"+host+"/"+user+"?sslmode=disable", "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	//Create a logger
	//logs anything that not an error.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//logs an error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var db, err = setUpDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() //can only close the connection where the function is called

	//passing the database connection to a model so it can handle other database operations
	app := &application{
		quotes: &postgresql.QuoteModel{
			DB: db,
		},
		errorLog: errorLog,
		infoLog: infoLog,
	}

	//create a custom web server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(), //return the multiplexer
		ErrorLog: errorLog, // initialize the standard error log with my own errorlog
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)

}
