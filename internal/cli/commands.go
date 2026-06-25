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
	if len(list) == 0 {
		fmt.Println("You are all done !!")
	} else {
		for i, v := range list {
			fmt.Printf("%s\n\t%d.%s", Dim, i, Reset)
			fmt.Printf("%s%s\t%s\n%s", Bold, Green, v.Title, Reset)
		}
	}
}

func ListCompleted(q *store.Queries, session string) {
	list, err := q.ListCompleted(context.Background(), session)
	if err != nil {
		log.Fatalln("Error fetching data from completed table: ", err)
	}
	if len(list) == 0 {
		fmt.Println("You have completed Nothing till !!")
	} else {
		for i, v := range list {
			fmt.Printf("%s\n\t%d.%s", Dim, i, Reset)
			fmt.Printf("%s%s\t%s\n%s", Bold, Green, v.Title, Reset)
		}
	}
}

func ListTrash(q *store.Queries, session string) {
	list, err := q.ListTrash(context.Background(), session)
	if err != nil {
		log.Fatalln("Error fetching data from completed table: ", err)
	}
	if len(list) == 0 {
		fmt.Println(Yellow, "You haven't removed anything yet !!", Reset)
	} else {
		for i, v := range list {
			fmt.Printf("%s\n\t%d.%s", Dim, i, Reset)
			fmt.Printf("%s%s\t%s\n%s", Bold, Green, v.Title, Reset)
		}
	}
}

func ListIgnored(q *store.Queries, session string) {
	list, err := q.ListIgnored(context.Background(), session)
	if err != nil {
		log.Fatalln("Error fetching data from ignored table: ", err)
	}
	if len(list) == 0 {
		fmt.Println(Yellow, "You haven't ignored anything yet !!", Reset)
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

	err = q.TrashTodoTransaction(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln("Error while transferign data from todo table to trash table, ", err)
	}
	err = q.DeleteFromTodos(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Cyan, id, " Has been succesuflly deleted", Reset)
}

func RemoveCompleted(q *store.Queries, id int, session string) {
	dbId, err := q.ListCompletedIDs(context.Background(), session)
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}

	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit completed list>", Reset)
	} else if len(dbId) == 0 {
		log.Println(Yellow, "You haven't Completed anything YET !!", Reset)
		os.Exit(0)
	}

	err = q.DeleteFromCompleted(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Cyan, id, " Has been succesuflly deleted", Reset)
}

func RemoveTrash(q *store.Queries, id int, session string) {
	dbId, err := q.ListTrashIDs(context.Background(), session)
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}

	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit trash list>", Reset)
	} else if len(dbId) == 0 {
		log.Println(Yellow, "You haven't Removed anything YET !!", Reset)
		os.Exit(0)
	}

	err = q.DeleteFromTrash(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Cyan, id, " Has been succesuflly deleted", Reset)
}

func RemoveIgnored(q *store.Queries, id int, session string) {
	dbId, err := q.ListIgnoredIDs(context.Background(), session)
	if err != nil {
		log.Fatalln(Red, "Error while retriving ID's from SQLite", err, Reset)
	}

	if len(dbId) < id {
		log.Fatalln(Red, "Provide correct id <use: doit trash list>", Reset)
	} else if len(dbId) == 0 {
		log.Println(Yellow, "You haven't Ignored anything YET !!", Reset)
		os.Exit(0)
	}

	err = q.DeleteFromIgnored(context.Background(), dbId[id])
	if err != nil {
		log.Fatalln(Red, "Error while deleting: ", err, Reset)
	}

	log.Println(Cyan, id, " Has been succesuflly deleted", Reset)
}

func MoveTrash(q *store.Queries, id int64) error {
	err := q.MoveTrashTo(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func MoveTrashToCompleted(q *store.Queries, id int64) error {
	err := q.MoveTrashToCompleted(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func MoveTrashToIgnored(q *store.Queries, id int64) error {
	err := q.MoveTrashToIgnored(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func ClearCompleted(q *store.Queries) {
	q.ClearCompleted(context.Background())
}
func ClearIgnored(q *store.Queries) {
	q.ClearIgnored(context.Background())
}
func ClearTrash(q *store.Queries) {
	q.ClearTrash(context.Background())
}
