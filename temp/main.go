package main

import (
	"fmt"
	"log"
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
	list := args[2:]
	newTaskList := make([]string, 0, len(args[2:]))
	// var newTaskList []string
	//
	for index, v := range list {
		if v == "," && list[index+1][0] != '-' {
			log.Println(list[i+1][0])
			commaCount++
		}
		if v != "," && v[0] != '-' {
			newTaskList = append(newTaskList, v)
			i++
		}
		log.Println("loop: ", index, " done")
	}

	log.Println(newTaskList)
	if i != commaCount+1 {
		os.Exit(106)
	}
	fmt.Printf("\nAdding:%d\ttaks to todoList\n", i)
	fmt.Printf("\ntodoList:%v,\tLen:%d,\tCap:%d", newTaskList, len(newTaskList), cap(newTaskList))

}
