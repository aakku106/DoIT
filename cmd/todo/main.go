package main

import (
	"log"

	"github.com/aakku106/DoIT/internal/db"
)

var DEBUG = true

func main() {
	if DEBUG {
		log.Println("Starting Program ...")
	}
	db.NewSQLite("./testi.db")
	if DEBUG {
		log.Println("XXXXXXXXXXX---------ENDING  Program ------XXXXXXXXXXXXX")
	}
}
