package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func doneTodo(q *store.Queries, args []string) {
	if len(args) < 3 || len(args) != 3 {
		fmt.Println(cli.Cyan, "Specify what have you done ? run exactly doit <d.done> id", cli.Reset)
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[2])
	if err != nil || id < 0 {
		fmt.Println(cli.Red, cli.Bold, "Enter valid id", err, cli.Reset)
	}
	cli.DoneTodo(q, id)
}
