package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

// TODO: Fix logic here
func clearTodo(q *store.Queries) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(cli.Red, "___This clears all tasks from Todo list. Tasks CANNOT be retrieved again!___", cli.Reset)
	fmt.Print(cli.Red, "---Are you sure you want to clear Todo (Y/N)? ", cli.Reset)

	// Read first input line
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	if len(input) == 0 {
		fmt.Println("Invalid selection.")
		return
	}

	// 'n' or 'N' cancels the operation
	if strings.ToLower(input) == "n" {
		fmt.Println("Operation cancelled.")
		return
	} else if input == "Y" { // Strictly checks for uppercase 'Y'
		fmt.Println(cli.Red, "CONFIRM CLEARING TODO LIST:", cli.Reset, "Type 'YeS NuKe ToDos' to confirm:")

		// Read second input line (with spaces)
		confirmInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation:", err)
			return
		}

		confirmInput = strings.TrimSpace(confirmInput)

		if confirmInput == "YeS NuKe ToDos" {
			cli.ClearTodos(q)
			fmt.Println("Todo list successfully cleared!")
			return
		}

		fmt.Println("CLEARING TODO LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between uppercase 'Y' and 'N'/'n'")
	}
}

func clearCompleted(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Completed list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Completed :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING TODO LIST:", cli.Reset, "YeS_NuKe / N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS_NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING COMPLETED LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between Y and N")
		os.Exit(1)
	}
}

func clearTrash(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Trash list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Trash :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING TODO LIST:", cli.Reset, "YeS_NuKe / N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS_NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING TRASH LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between Y and N")
		os.Exit(1)
	}
}

func clearIgnored(q *store.Queries) {
	fmt.Println(cli.Red, "___This Clear all Tasks From Ignored list, Task Cannot be Retrived Again !!!___", cli.Reset)
	fmt.Println(cli.Red, "---Are you sure You want To Clear Ignored :", cli.Reset, "Y/N")
	var a rune
	fmt.Scanf("%c", &a)
	if unicode.ToLower(a) == 'n' {
		os.Exit(0)
	} else if a == 'Y' {
		fmt.Println(cli.Red, "CONFIRN CLEARING TODO LIST:", cli.Reset, "YeS_NuKe / N")
		var b string
		fmt.Scanf("%s", &b)
		if b == "YeS_NuKe" {
			cli.ClearCompleted(q)
		}
		fmt.Println("CLEARING IGNORED LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between Y and N")
		os.Exit(1)
	}
}
