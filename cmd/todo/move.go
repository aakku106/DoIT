package main

import (
	"context"
	"log"
	"strconv"

	"github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

func handleTrashMove(q *store.Queries, args []string) {
	id, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		log.Fatalln(cli.Red, "Enter valid id ", err, cli.Reset)
	}
	if len(args) == 4 {
		moveTrash(q, id)
	} else if len(args) == 5 {
		switch args[4] {
		case "completed", "c":
			moveTrashToCompleted(q, id)
		case "ignored", "i":
			moveTrashToIgnored(q, id)
		}
	}
}

func moveTrash(q *store.Queries, id int64) {
	q.MoveTrashTo(context.Background(), id)
}
func moveTrashToCompleted(q *store.Queries, id int64) {
	q.MoveTrashToCompleted(context.Background(), id)
}
func moveTrashToIgnored(q *store.Queries, id int64) {
	q.MoveCompletedToIgnored(context.Background(), id)
}

func handleComletedMove(q *store.Queries, args []string) {
	id, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		log.Fatalln(cli.Red, "Enter valid id ", err, cli.Reset)
	}
	if len(args) == 4 {
		moveComepeted(q, id)
	} else if len(args) == 5 {
		moveTrash(q, id)
	}
}

func moveComepeted(q *store.Queries, id int64)          {}
func moveComepetedToTrash(q *store.Queries, id int64)   {}
func moveComepetedToIgnored(q *store.Queries, id int64) {}

func handleIgnoredMove(q *store.Queries, args []string) {}

func moveIgnored(q *store.Queries, id int64)          {}
func moveIgnoredToTrash(q *store.Queries, id int64)   {}
func moveIgnoredCompleted(q *store.Queries, id int64) {}
