package cli

import (
	"fmt"

	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(q *store.Queries, title string) {
	fmt.Println("Add called got:", q, title)
}
func ListTodos(q *store.Queries)              {
	fmt.Println("Listing TODO:")
	q.ListTodos()
}
func CompleteTodo(q *store.Queries, id int64) {}
