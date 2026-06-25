package main

import (
	"context"
	"log"
	"strconv"

func moveTrash(q *store.Queries, args []string)            {}
func moveTrashToCompleted(q *store.Queries, args []string) {}
func moveTrashToIgnored(q *store.Queries, args []string)   {}

func moveComepeted(q *store.Queries, args []string)          {}
func moveComepetedToTrash(q *store.Queries, args []string)   {}
func moveComepetedToIgnored(q *store.Queries, args []string) {}

func moveIgnored(q *store.Queries, args []string)          {}
func moveIgnoredToTrash(q *store.Queries, args []string)   {}
func moveIgnoredCompleted(q *store.Queries, args []string) {}
