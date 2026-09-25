package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(cli.Red, "___This clears all tasks from Completed list. Tasks CANNOT be retrieved again!___", cli.Reset)
	fmt.Print(cli.Red, "---Are you sure you want to clear Completed (Y/N)? ", cli.Reset)

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
		fmt.Println(cli.Red, "CONFIRM CLEARING COMPLETED LIST:", cli.Reset, "Type 'YeS NuKe CoMpleteD' to confirm:")

		// Read second input line (with spaces)
		confirmInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation:", err)
			return
		}

		confirmInput = strings.TrimSpace(confirmInput)

		if confirmInput == "YeS NuKe CoMpleteD" {
			cli.ClearCompleted(q)
			fmt.Println("Completed list successfully cleared!")
			return
		}

		fmt.Println("CLEARING COMPLETED LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between uppercase 'Y' and 'N'/'n'")
	}
}

func clearTrash(q *store.Queries) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(cli.Red, "___This clears all tasks from Trash list. Tasks CANNOT be retrieved again!___", cli.Reset)
	fmt.Print(cli.Red, "---Are you sure you want to clear Trash (Y/N)? ", cli.Reset)

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
		fmt.Println(cli.Red, "CONFIRM CLEARING TRASH LIST:", cli.Reset, "Type 'YeS NuKe TrAsH' to confirm:")

		// Read second input line (with spaces)
		confirmInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation:", err)
			return
		}

		confirmInput = strings.TrimSpace(confirmInput)

		if confirmInput == "YeS NuKe TrAsH" {
			cli.ClearTrash(q)
			fmt.Println("Trash list successfully cleared!")
			return
		}

		fmt.Println("CLEARING TRASH LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between uppercase 'Y' and 'N'/'n'")
	}
}

func clearIgnored(q *store.Queries) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(cli.Red, "___This clears all tasks from Ignored list. Tasks CANNOT be retrieved again!___", cli.Reset)
	fmt.Print(cli.Red, "---Are you sure you want to clear Ignored (Y/N)? ", cli.Reset)

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
		fmt.Println(cli.Red, "CONFIRM CLEARING IGNORED LIST:", cli.Reset, "Type 'YeS NuKe IgNoreD' to confirm:")

		// Read second input line (with spaces)
		confirmInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation:", err)
			return
		}

		confirmInput = strings.TrimSpace(confirmInput)

		if confirmInput == "YeS NuKe IgNoreD" {
			cli.ClearIgnored(q)
			fmt.Println("Ignored list successfully cleared!")
			return
		}

		fmt.Println("CLEARING IGNORED LIST ABORTED !!!")
	} else {
		fmt.Println("You were supposed to select between uppercase 'Y' and 'N'/'n'")
	}
}
