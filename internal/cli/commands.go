package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(query *store.Queries, title string) {
	fmt.Printf("\t%s\n", title)
	param := store.CreateTodoParams{
		Title:     title,
		Session:   "todo",
		ExpiresAt: sql.NullTime{Valid: false},
	}
	query.CreateTodo(context.Background(), param)
}
func ListTodos() {
	fmt.Println("Listing TODO:")
}
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo()                         {}
