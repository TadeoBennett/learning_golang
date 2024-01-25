package main


import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" //third party package
)

//provide the credentials for our database
const (
	host = "localhost"
	port = 5432
	user = "quotebox"
	password = "tadeo2002"
	dbname = "quotebox"
)


//dsn: data source name
func main(){
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close() //needs to be closed to free space 

	err = db.Ping() //Test our connection
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println("Connection status: success")
}

	//NOTE: Go only gives us the resource if we asked for it. 
	//Go does not allocate space for something unless there is code for it


