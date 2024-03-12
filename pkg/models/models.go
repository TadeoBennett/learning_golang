//need to give the package a main (same as the directory it is in)
package models

import(
	"time"
)


//a struct to hold a quote
type Quote struct {
	Quotation_id   int
	Created_at time.Time
	Author_name    string
	Category       string
	Body          string
}


