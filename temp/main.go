package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Printf("\ntodoList:%v,\tLen:%d,\tCap:%d\n", val, len(val), cap(val))

	t := Task{}
	if val, err := t.sanitizeTask(val); err != nil {
		log.Panic(err)
	} else {
		fmt.Println("value ", val, "cap ", cap(val), "len ", len(val))
	}
}

func extractMultipleTasks(args []string) ([][]string, error) {

	rawArgs := args[2:]

	// Bucket for currentTask and Group of Current Task
	taskGroups := make([][]string, 0, len(rawArgs)/2)
	currentGroup := make([]string, 0, 2) // cap=2 cause, for now we only have 2 possible entrires{task,time}

	// Group arguments delimited by ","
	for _, val := range rawArgs {
		if val == "," {
			if len(currentGroup) == 0 {
				return nil, fmt.Errorf("Comma Before any task")
			}
			taskGroups = append(taskGroups, currentGroup)
			currentGroup = currentGroup[:0] // I avoided using []string{}, cause it escapes to heap
		} else {
			currentGroup = append(currentGroup, val)
		}
	}

	if len(currentGroup) == 0 {
		return nil, fmt.Errorf("comma in last without any task.")
	}
	// pushing final task in group bucket
	taskGroups = append(taskGroups, currentGroup)

	return taskGroups, nil
}

// Verify each task from group bucket, input [][]string and clean any given time,
// For now we only check for title and time.
// Title shall never start with , or -  everythign else is valid
// Title shall never end with ,, everything else is valid
func (t *Task) sanitizeTask(v [][]string) ([]Task, error) {
	task := make([]Task, 0, len(v))

	for _, value := range v {
		_title := value[0]
		_time := value[1]
		{
			fmt.Printf("\nTitle:%s\ttime:%s", _title, _time)
		}
		if strings.HasPrefix(_title, "-") {
			return nil, fmt.Errorf("Expected task name, got flag '%s'\n", _title)
		}
		if len(_time) == 0 {
			fmt.Println("DeadLine time not assigned")
		}
		if len(_time) != 0 && !strings.HasPrefix(_time, "-t=") {
			return nil, fmt.Errorf("Dead line time should be given as -t=")
		} else {
			_time = strings.TrimPrefix(_time, "-t=")
		}

		task = append(task, Task{_title, _time})
	}
	return task, nil
}
