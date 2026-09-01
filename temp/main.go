package main

import (
	"fmt"
	"os"
)

func main() {

	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fmt.Println(pwd)
	}

	if err := os.Mkdir("CAt", 0755); err != nil {
		panic(err)
	}

}
