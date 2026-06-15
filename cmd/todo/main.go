package main

import (
	"fmt"
	"os"

	call "github.com/aakku106/DoIT/internal/cli"
)

func main() {
	args := os.Args
	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Specify what to add !")
			os.Exit(1)
		}
		fmt.Print("Adding:")
		add(args)

	case "list":
		fmt.Println("Listing data of todo")
		list()

	case "done":
		fmt.Println(args[2], "task completed")
		if len(args) < 3 {
			fmt.Println("Specify what have you done ?")
			os.Exit(1)
		}
		done(args)

	case "remove":
		if len(args) < 3 {
			fmt.Println("Specify what to Remove ?")
			os.Exit(1)
		}
		fmt.Println("Removing: ", args[3])
		remove(args)

	default:
		fmt.Println("Session call")
		if len(args) <= 2 {
			fmt.Println("Bro what >???<")
			os.Exit(1)
		}
		sessionCall(args)
	}
}

func add(args []string) {
	call.AddTodo(args[2])
}
func list() {
	call.ListTodos()
}
func done(args []string) {}
func remove(args []string) {
	fmt.Println("***---Are you sure You want To remove: ", args[2])
}

func sessionCall(args []string) {}
