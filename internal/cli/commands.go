package cli

import (
	"fmt"
	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(query *store.Queries, title string) {
	fmt.Printf("\t%s\n", title)

}
func ListTodos() {
	fmt.Println("Listing TODO:")
}
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo()                         {}
