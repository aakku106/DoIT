package cli

import (
	"fmt"
	"time"

	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(title string, time time.Time) {
	fmt.Printf("%s in", time)
	fmt.Println("Called on", time)
	fmt.Println("Local: ", time.Local())
}
func ListTodos(q *store.Queries) {
	fmt.Println("Listing TODO:")
}
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo()                         {}
