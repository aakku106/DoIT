package main

import (
	"fmt"

	"github.com/aakku106/DoIT/internal/store"

	"github.com/aakku106/DoIT/internal/cli"
)

func listTodo(q *store.Queries, args []string) {
	fmt.Print("\033[H\033[2J")
	if len(args) == 2 {
		fmt.Println("Listing data from 'todo'")
		cli.ListTodos(q)
	}
}

func listCompleted(q *store.Queries, args []string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Listing data from 'completed'")
	cli.ListCompleted(q, "todo")
}

func listTrash(q *store.Queries, args []string) {
	fmt.Print("\033[H\033[2J")
	if len(args) == 3 {
		fmt.Println("Listing data from 'trash'")
		cli.ListTrash(q, "todo")
	}
}

func listIgnored(q *store.Queries, args []string) {
	fmt.Print("\033[H\033[2J")
	if len(args) == 3 {
		fmt.Println("Listing data from 'ignored'")
		cli.ListIgnored(q, "todo")
	}
}
