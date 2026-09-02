package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// fmt.Printf("Args: %v \nLength of arg: %d ", args, len(args))
	if args[1] != "add" {
		os.Exit(106)
	}
	if len(args) < 3 {
		os.Exit(106)
	}

	fmt.Println("Addign Multiple tasks to todo list")

	i := 0
	commaCount := 0
	var newTaskList []string

	for _, v := range args[2:] {
		if v == "," {
			commaCount++
		}

		if v != "," {
			if v[0] != '-' {
				newTaskList = append(newTaskList, v)
				i++
			}
		}

	}
	if i != commaCount+1 {
		os.Exit(106)
	}
	fmt.Printf("\nAdding:%d\ttaks to todoList\n", i)
	fmt.Printf("\ntodoList:%v,\tLen:%d,\tCap:%d", newTaskList, len(newTaskList), cap(newTaskList))

}
