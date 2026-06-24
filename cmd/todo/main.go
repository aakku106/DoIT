package main

import (
	"fmt"
	"os"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/db"
	"github.com/aakku106/DoIT/internal/store"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Bro what >???<")
		os.Exit(1)
	}

	db, err := db.NewSQLite()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer db.Close()

	query := store.New(db)

	switch args[1] {
	case "add":
		add(query, args)

	case "list":
		listTodo(query, args)

	case "done":
		doneTodo(query, args)

	case "remove":
		removeTodo(query, args)

	case "completed":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !!")
		}
		switch args[2] {
		case "list":
			listCompleted(query, args)

		case "remove":
			removeCompleted(query, args)

		case "nuke":
			clearCompleted(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
		}

	case "trash":
		switch args[2] {
		case "list":
			listTrash(query, args)

		case "remove":
			removeTrash(query, args)

		case "nuke":
			clearTrash(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
		}

	default:
		if len(args) <= 2 {
			fmt.Println("Bro what >???<")
			os.Exit(1)
		}
		fmt.Println("Session cli")
		sessionCall(args)
	}
}

func sessionCall(args []string) {}
