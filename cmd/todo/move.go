package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func handleTrashMove(q *store.Queries, args []string) {
	id, err := strconv.ParseInt(args[3], 10, 64)

	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}

	if len(args) == 4 {
		if err := q.MoveTrashTo(context.Background(), id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		}
	} else if len(args) == 5 {
		switch args[4] {
		case "completed", "c":
			if err := q.MoveTrashToCompleted(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}

		case "ignored", "i":
			if err := q.MoveTrashToIgnored(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}
		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}

func handleComletedMove(q *store.Queries, args []string) {
	id, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}
	if len(args) == 4 {
		if err := q.MoveCompletedTo(context.Background(), id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		}
	} else if len(args) == 5 {
		switch args[4] {
		case "trash", "t":
			if err := q.MoveCompletedToTrash(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}
		case "ignored", "i":
			if err := q.MoveCompletedToIgnored(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}
		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}

func handleIgnoredMove(q *store.Queries, args []string) {
	id, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		fmt.Println(cli.Red, "Enter valid id ", err, cli.Reset)
		os.Exit(1)
	}
	if len(args) == 4 {
		if err := q.MoveIgnoredTo(context.Background(), id); err != nil {
			log.Fatalln(cli.Red, err, cli.Reset)
		}
	} else if len(args) == 5 {
		switch args[4] {
		case "trash", "t":
			if err := q.MoveIgnoredToTrash(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}
		case "completed", "c":
			if err := q.MoveIgnoredToCompleted(context.Background(), id); err != nil {
				log.Fatalln(cli.Red, err, cli.Reset)
			}
		}
	} else {
		fmt.Println("Run exactly: ", cli.Bold, " doit <t/trash> id target/Optional/")
		os.Exit(1)
	}
}
