package main

import (
	"fmt"
	"os"

	call "github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func add(q *store.Queries, args []string) {
	if len(args) < 3 {
		fmt.Println(call.Cyan, "Specify what to add !, run exactly doit <a/add> your task", call.Reset)
		os.Exit(1)
	}
	fmt.Print("Adding:")
	if len(args) == 3 {
		call.AddTodo(q, args[2])
	} else if len(args) > 3 {
		// TODO:Work is going on on temp/main.go
	}
}
