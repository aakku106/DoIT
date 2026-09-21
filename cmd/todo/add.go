package main

import (
	"fmt"
	"log"
	//"os"
	"strings"

	call "github.com/aakku106/DoIT/internal/cli"
	"github.com/aakku106/DoIT/internal/store"
)

type Task struct {
	Title string
	Time  string
}

const DeadMissingWarningMessage bool = false

func add(q *store.Queries, args []string) {
	// if len(args) < 3 {
	// 	fmt.Println(call.Cyan, "Specify what to add !, run exactly doit <a/add> your task", call.Reset)
	// 	os.Exit(1)
	// }
	if len(args) == 3 {
		fmt.Println("Adding:")
		call.AddTodo(q, args[2])
	} else if len(args) > 3 {

		fmt.Println("Adding Multiple Tasks:")
		val, err := extractMultipleTasks(args)
		if err != nil {
			log.Panic(err)
		}
		t := Task{}
		if value, err := t.sanitizeTasks(val); err != nil {
			log.Panic(err)
		} else {
			for _, v := range value {
				call.AddTodo(q, v.Title)
			}
		}
	}
}

// Extract Tasks and DeadLine from given to task or group of tasks.
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
			currentGroup = []string{}
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
func (t *Task) sanitizeTasks(v [][]string) ([]Task, error) {
	task := make([]Task, 0, len(v))

	for _, value := range v {

		_title := value[0]
		if strings.HasPrefix(_title, "-") && strings.HasPrefix(_title, ",") {
			return nil, fmt.Errorf("Task Name, shouldn't have presiding flag '-' or comma ',', but got '%s'\n", _title)
		}

		var _time string
		if len(value) > 1 {
			_time = value[1]
		}

		if len(_time) == 0 && DeadMissingWarningMessage {
			fmt.Println("DeadLine time not assigned, you can assigne deadline on any task 'doit add TaskName -t=2h , \"Another Task\" -t=1mo '")
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
