package cli

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo(q *store.Queries, id int) {
	dbId, err := q.ListTodoIDs(context.Background(), "todo")
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}
	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit list>", Reset)
	}
	q.DeleteTodo(context.Background(), dbId[id])
}
