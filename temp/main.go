package main

import (
	"fmt"
	"log"
	"os"
)

type Task struct {
	Title string
	Time  string
}

func main() {

	args := os.Args

	if args[1] != "add" {
		os.Exit(106)
	}
	if len(args) < 3 {
		os.Exit(106)
	}

	fmt.Println("Addign Multiple tasks to todo list")

	val, err := extractMultipleTasks(args)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(val)
	fmt.Printf("\ntodoList:%v,\tLen:%d,\tCap:%d\n", val, len(val), cap(val))
}

func extractMultipleTasks(args []string) ([][]string, error) {

	rawArgs := args[2:]

	// Bucked for currentTask and Group of Current Task
	taskGroups := make([][]string, 0, len(rawArgs)/2)
	currentGroup := make([]string, 0, 2) // cap=2 cause, for now we only have 2 possible entrires{task,time}

	for _, val := range rawArgs {
		if val == "," {
			if len(currentGroup) == 0 {
				return nil, fmt.Errorf("Comma Before any task")
			}
			taskGroups = append(taskGroups, currentGroup)
			currentGroup = []string{}
		} else {
			currentGroup = append(currentGroup, val)
		}
	}
	if len(currentGroup) == 0 {
		return nil, fmt.Errorf("comma in last without any task.")
	}
	taskGroups = append(taskGroups, currentGroup)

	fmt.Printf("\ntodoList:%v,\tLen:%d,\tCap:%d", taskGroups, len(taskGroups), cap(taskGroups))

}
