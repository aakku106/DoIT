package main

import (
	"github.com/aakku106/DoIT/internal/db"
	"log"
)

func main() {
	qry, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("qry: %v\n", qry)
}
