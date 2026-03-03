package main

import (
	"fmt"
	"log"

	"github.com/aakku106/DoIT/internal/db"
)

var DEBUG = true

func main() {
	if DEBUG {
		log.Println("Starting Program ...")
	}
	qry, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("qry: %v\n", qry)
	if DEBUG {
		log.Println("XXXXXXXXXXX---------ENDING  Program ------XXXXXXXXXXXXX")
	}
}
