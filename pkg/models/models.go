// need to give the package a main (same as the directory it is in)
package models

import (
	"errors"
	"time"
)

//creating a new error for this model
var ErrRecordNotFound = errors.New("models: no matching record found")

//a struct to hold a quote
type Quote struct {
	Quotation_id   int
	Created_at time.Time
	Author_name    string
	Category       string
	Body          string
}


