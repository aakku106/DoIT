package cli

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(query *store.Queries, title string) {
	fmt.Printf("\t%s\n", title)
	param := store.CreateTodoParams{
		Title:     title,
		Session:   "todo",
		ExpiresAt: sql.NullTime{Valid: false},
	}
	added, err := query.CreateTodo(context.Background(), param)
	if err != nil {
		log.Fatalln("Error while adding to DB: ", err)
	}
	fmt.Println("Succesfully Added ToDo")
	fmt.Println("ID:\t\t", Bold, Green, added.ID, Reset)
	fmt.Println("ToDo Task:\t", Bold, Green, added.Title, Reset)
	fmt.Println("Created At:\t", Green, added.CreatedAt.Local(), Reset)
}
func ListTodos(q *store.Queries) {
	list, err := q.ListTodos(context.Background(), "todo")
	if err != nil {
		log.Fatalln("Error while listing todo data from DB: ", err)
	}
	if list == nil {
		fmt.Println("You are all done !!")
	} else {
		for i, v := range list {
			fmt.Printf("%s\n\t%d.%s", Dim, i, Reset)
			fmt.Printf("%s%s\t%s\n%s", Bold, Green, v.Title, Reset)
		}
	}
}

func DoneTodo(q *store.Queries, id int) {
	dbId, err := q.ListTodoIDs(context.Background(), "todo")
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}

	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit list>", Reset)
	} else if len(dbId) == 0 {
		log.Println(Yellow, "You have no to do , You are all Done !!", Reset)
		os.Exit(0)
	}

	err = q.CompleteTodoTransaction(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}
	err = q.DeleteFromTodos(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Green, " Congrats on Completoin of task", id, Reset)
}

func RemoveTodo(q *store.Queries, id int) {
	dbId, err := q.ListTodoIDs(context.Background(), "todo")
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}

	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit list>", Reset)
	} else if len(dbId) == 0 {
		log.Println(Yellow, "You have no to do , You are all Done !!", Reset)
		os.Exit(0)
	}

	err = q.DeleteTodo(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Cyan, id, " Has been succesuflly deleted", Reset)
}
