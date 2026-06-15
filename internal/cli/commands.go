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
	fmt.Println("ID:\t", added.ID)
	fmt.Println("ToDo Task:\t", added.Title)
	fmt.Println("Created At:\t", added.CreatedAt.Local())
}
func ListTodos() {
	fmt.Println("Listing TODO:")
}
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo()                         {}
