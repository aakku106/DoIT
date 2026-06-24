package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func removeTodo(q *store.Queries, args []string) {
	if len(args) < 3 {
		fmt.Println(cli.Cyan, "Specify what to Remove ?'provide task Id <use doit list>'", cli.Reset)
		os.Exit(1)
	}

	fmt.Println("---Are you sure You want To remove:", cli.Red, args[2], cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else { // we do need this else, cause lest say if user gave "wotver except N/n" it will still delete args[2] id from db
		id, err := strconv.Atoi(args[2])
		if err != nil || id < 0 {
			fmt.Println(cli.Red, cli.Bold, "Enter valid id", err, cli.Reset)
		}
		fmt.Println("Trying to Remove: ", cli.Bold, cli.Red, id, cli.Reset)
		cli.RemoveTodo(q, id)
	}
}
