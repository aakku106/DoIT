package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// fmt.Printf("Args: %v \nLength of arg: %d ", args, len(args))
	if args[1] == "add" {
		if len(args) > 3 {
			fmt.Println("Addign Multiple tasks to todo list")
			i := 1 // just to know hou many times loop ran
			for _, v := range args[2:] {
				if v != "," {
					fmt.Printf("\nIndex:%d\tvalue: %s\n", i, v)
					i++
				}
			}
		}
	}
}
