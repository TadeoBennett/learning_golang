package main

import (
	"tadeobennett.net/quotation/pkg/models"
)

type templateData struct {
	//a slice of pointers to Quote struct
	Quotes []*models.Quote
}


