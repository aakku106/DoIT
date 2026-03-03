package main

import (
	"context"
	"log"

	"github.com/aakku106/DoIT/internal/db"
	"github.com/aakku106/DoIT/internal/store"
)

func main() {
	dbConn, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	log.Println("DB opened")
	queries := store.New(dbConn)

	todo, err := queries.CreateTodo(context.Background(), store.CreateTodoParams{
		Title:   "Learn sqlc properly",
		Session: "todo",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted: %+v\n", todo)

}
