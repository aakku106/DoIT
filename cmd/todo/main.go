package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	call "github.com/aakku106/DoIT/internal/cli"
	create "github.com/aakku106/DoIT/internal/db"
	"github.com/aakku106/DoIT/internal/store"
)

func main() {
	args := os.Args
	db, err := create.NewSQLite()
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
			fmt.Println(call.Red, "---Are you sure You want To Clear Completed :", call.Reset, "Y/N")
			var a rune
			fmt.Scanf("%c", &a)
			if unicode.ToLower(a) == 'n' {
				os.Exit(0)
			} else {
				call.ClearCompleted(query)
			}

		default:
			fmt.Println(call.Cyan, "Umm wot ??", call.Reset)
		}

	case "trash":
		switch args[2] {
		case "list":
			listTrash(query, args)

		case "remove":
			removeTrash(query, args)
		case "nuke":
			fmt.Println(call.Red, "---Are you sure You want To Clear Completed :", call.Reset, "Y/N")
			var a rune
			fmt.Scanf("%c", &a)
			if unicode.ToLower(a) == 'n' {
				os.Exit(0)
			} else {
				call.ClearTrash(query)
			}

		default:
			fmt.Println(call.Cyan, "Umm wot ??", call.Reset)
		}

	default:
		if len(args) <= 2 {
			fmt.Println("Bro what >???<")
			os.Exit(1)
		}
		fmt.Println("Session call")
		sessionCall(args)
	}
}

func sessionCall(args []string) {}
