package main

import (
	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/db"
	"github.com/aakku106/DoIT/internal/store"
	"log"
	"os"
)

func main() {
	dbConn, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	log.Println("DB opened")
	queries := store.New(dbConn)
	log.Println(queries)

	args := os.Args

	switch args[1] {
	case "add":
		cli.AddTodo(queries, args[1])
	case "list":
		cli.ListTodos(queries)
	case "done":
		cli.CompleteTodo(queries, args[2])
	default:
		log.Println("unknown command")
	}

}
