// Copyright (C) 2026 Adarasha Gaihre (aakku) <aakku@adarashagaihre.com.np>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"fmt"
	"os"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/db"
	"github.com/aakku106/DoIT/internal/store"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Bro what >???<")
		os.Exit(1)
	}

	db, err := db.NewSQLite()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer db.Close()

	query := store.New(db)

	switch args[1] {
	case "add", "a":

		add(query, args)

	case "list", "ls":
		listTodo(query, args)

	case "done", "d":
		doneTodo(query, args)

	case "remove", "rm":
		removeTodo(query, args)

	case "completed", "c":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !! run, doit <c/completed> [list | remove | nuke] ")
			os.Exit(1)
		}
		switch args[2] {
		case "list", "ls":
			listCompleted(query, args)

		case "remove", "rm":
			removeCompleted(query, args)

		case "nuke", "n":
			clearCompleted(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
			os.Exit(1)
		}

	case "trash", "t":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !! run, doit <t/trash> [list | remove | nuke] ")
			os.Exit(1)
		}
		switch args[2] {
		case "list", "ls":
			listTrash(query, args)

		case "remove", "rm":
			removeTrash(query, args)

		case "nuke", "n":
			clearTrash(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
			os.Exit(1)
		}

	case "ignored", "i":
		if len(args) < 3 {
			fmt.Println("Umm WOt Broo !! run, doit <i/ignored> [list | remove | nuke] ")
			os.Exit(1)
		}
		switch args[2] {
		case "list", "ls":
			listIgnored(query, args)

		case "remove", "rm":
			removeIgnored(query, args)

		case "nuke", "n":
			clearIgnored(query)

		default:
			fmt.Println(cli.Cyan, "Umm wot ??", cli.Reset)
			os.Exit(1)
		}

	case "move", "mv":

		if len(args) < 4 {
			fmt.Println("Umm Wot Broo !! run, doit <mv/move> [<t/trash> | <i/ignored> | <c/completed>] id  [<t/trash> | <i/ignored> | <c/completed>]")
			os.Exit(1)
		}

		switch args[2] {
		case "trash", "t":
			handleTrashMove(query, args)
		case "completed", "c":
			handleComletedMove(query, args)
		case "ignored", "i":
			handleIgnoredMove(query, args)
		default:
			fmt.Println(cli.Cyan, "move commad only work for trash,ignored,completed tables", cli.Reset)
		}

	default:
		if len(args) <= 2 {
			fmt.Println("Bro what >???<")
			os.Exit(1)
		}
		fmt.Println("Session cli")
		sessionCall(args)
	}
}
