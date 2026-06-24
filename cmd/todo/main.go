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
	case "add", "a":

		add(query, args)

	case "list", "ls":
		listTodo(query, args)

	case "done", "d":
		doneTodo(query, args)

	case "remove", "rm":
		removeTodo(query, args)

	case "completed", "c":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !! run, doit completed [list | remove | nuke] ")
			os.Exit(1)
		}
		switch args[2] {
		case "list", "ls":
			listCompleted(query, args)

		case "remove", "rm":
			removeCompleted(query, args)

		case "nuke", "n":
			clearCompleted(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
			os.Exit(1)
		}

	case "trash", "t":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !! run, doit trash [list | remove | nuke] ")
			os.Exit(1)
		}
		switch args[2] {
		case "list", "ls":
			listTrash(query, args)

		case "remove", "rm":
			removeTrash(query, args)

		case "nuke", "n":
			clearTrash(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
			os.Exit(1)
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
