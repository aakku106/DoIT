package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func handleTrashMove(q *store.Queries, args []string) {
	id, err := strconv.Atoi(args[3])

	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}

	if len(args) == 4 {
		if err := cli.MoveTrash(q, id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		} else {
			fmt.Println(cli.Yellow, "Task Moved from 'Trash' to 'Todos")
		}

	} else if len(args) == 5 {
		switch args[4] {
		case "completed", "c":
			if err := cli.MoveTrashToCompleted(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Trash' to 'Completed")
			}

		case "ignored", "i":
			if err := cli.MoveTrashToIgnored(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Trash' to 'Ignored")
			}

		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}

func handleComletedMove(q *store.Queries, args []string) {
	id, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}
	if len(args) == 4 {
		if err := cli.MoveCompleted(q, id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		} else {
			fmt.Println(cli.Yellow, "Task Moved from 'Completed' to 'Toodos")
		}
	} else if len(args) == 5 {
		switch args[4] {
		case "trash", "t":
			if err := cli.MoveCompletedToTrash(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Completed' to 'Trash'")
			}

		case "ignored", "i":
			if err := cli.MoveCompletedToIgnored(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Completed' to 'Ignored'")
			}

		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}

func handleIgnoredMove(q *store.Queries, args []string) {
	id, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}
	if len(args) == 4 {
		if err := cli.MoveIgnored(q, id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		} else {
			fmt.Println(cli.Yellow, "Task Moved from 'Ignored' to 'Todos'")
		}

	} else if len(args) == 5 {
		switch args[4] {
		case "trash", "t":
			if err := cli.MoveIgnoredToTrash(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Ignored' to 'Trash'")
			}

		case "completed", "c":
			if err := cli.MoveIgnoredToCompleted(q, id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			} else {
				fmt.Println(cli.Yellow, "Task Moved from 'Ignored' to 'Completed'")
			}
		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}
