package main

import (
	"os"
)

func main() {
	if err := os.Mkdir("CAt", 0755); err != nil {
		panic(err)
	}
}
