package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func clearCompleted(q *store.Queries) {
	fmt.Println(cli.Red, "---Are you sure You want To Clear Completed :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else {
		cli.ClearCompleted(q)
	}
}

func clearTrash(q *store.Queries) {
	fmt.Println(cli.Red, "---Are you sure You want To Clear Completed :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else {
		cli.ClearTrash(q)
	}
}
