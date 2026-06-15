package cli

import (
	"fmt"
	"github.com/aakku106/DoIT/internal/store"
)

func AddTodo(title string) {
	fmt.Printf("%s in", title)

}
func ListTodos() {
	fmt.Println("Listing TODO:")
}
func DoneTodo(q *store.Queries, id int64) {}
func RemoveTodo()                         {}
