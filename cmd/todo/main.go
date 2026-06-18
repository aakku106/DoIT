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
		if len(args) < 3 {
			fmt.Println(call.Cyan, "Specify what to add !", call.Reset)
			os.Exit(1)
		}
		fmt.Print("Adding:")
		call.AddTodo(query, args[2])

	case "list":
		fmt.Print("\033[H\033[2J")
		if len(args) == 2 {
			fmt.Println("Listing data from 'todo'")
			call.ListTodos(query)
		}

	case "done":
		fmt.Println(args[2], "task completed")
		if len(args) < 3 {
			fmt.Println(call.Cyan, "Specify what have you done ?", call.Reset)
			os.Exit(1)
		}

	case "remove":
		if len(args) < 3 {
			fmt.Println(call.Cyan, "Specify what to Remove ?'provide task Id <use doit list>'", call.Reset)
			os.Exit(1)
		}

		fmt.Println("---Are you sure You want To remove:", call.Red, args[2], call.Reset, "Y/N")
		var a rune
		fmt.Scanf("%c", &a)
		if unicode.ToLower(a) == 'n' {
			os.Exit(0)
		} else { // we do need this else, cause lest say if user gave "wotver except N/n" it will still delete args[2] id from db
			id, err := strconv.Atoi(args[2])
			if err != nil || id < 0 {
				fmt.Println(call.Red, call.Bold, "Enter valid id", err, call.Reset)
			}
			fmt.Println("Trying to Remove: ", call.Bold, call.Red, id, call.Reset)
			call.RemoveTodo(query, id)
		}

	default:
		fmt.Println("Session call")
		if len(args) <= 2 {
			fmt.Println("Bro what >???<")
			os.Exit(1)
		}
		sessionCall(args)
	}
}

func sessionCall(args []string) {}
