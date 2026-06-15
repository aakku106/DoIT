package main

import (
	"fmt"
	"os"

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
			fmt.Println("Specify what to add !")
			os.Exit(1)
		}
		fmt.Print("Adding:")
		call.AddTodo(query, args[2])

	case "list":
		fmt.Println("Listing data of todo")

	case "done":
		fmt.Println(args[2], "task completed")
		if len(args) < 3 {
			fmt.Println("Specify what have you done ?")
			os.Exit(1)
		}

	case "remove":
		if len(args) < 3 {
			fmt.Println("Specify what to Remove ?")
			os.Exit(1)
		}

		fmt.Println("***---Are you sure You want To remove: ", args[2])
		fmt.Println("Removing: ", args[3])

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
