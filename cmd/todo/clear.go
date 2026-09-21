package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func clearTodo(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Todo list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Todo :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING TODO LIST:", cli.Reset, "YeS NuKe/N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS NuKe" {
			cli.ClearTodos(q)
		}
		fmt.Println("CLEARING TODO LIST ABORTED !!!")
	}
	fmt.Println("You were supposed to select between Y and N")
	os.Exit(1)
}

func clearCompleted(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Completed list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Completed :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING COMPLETED LIST:", cli.Reset, "YeS NuKe/N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING COMPLETED LIST ABORTED !!!")
	}
	fmt.Println("You were supposed to select between Y and N")
	os.Exit(1)
}

func clearTrash(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Trash list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Trash :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING TRASH LIST:", cli.Reset, "YeS NuKe/N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING TRASH LIST ABORTED !!!")
	}
	fmt.Println("You were supposed to select between Y and N")
	os.Exit(1)
}

func clearIgnored(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Ignored list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Ignored :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING IGNORED LIST:", cli.Reset, "YeS NuKe/N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING IGNORED LIST ABORTED !!!")
	}
	fmt.Println("You were supposed to select between Y and N")
	os.Exit(1)
}
