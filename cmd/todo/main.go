package main

import (
	"github.com/aakku106/DoIT/internal/db"
	"log"
)

func main() {
	dbConn, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	log.Println("DB opened")
}
