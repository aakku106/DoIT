package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getpagesize())
	// Lest re-look pointers
	a := 3
	b := &a
	fmt.Println("address: ", b, " Value: ", *b)
	// b = 3 // this cannot be dont, we cant change pointers
	// But we can
	*b = 106
	fmt.Println("address: ", b, " Value: ", *b)
	/*
		address:  0x716cca086038  Value:  3
		address:  0x716cca086038  Value:  106
	*/
	// See it shows same address, so that means we can change value that a pointer points at, put cant change the value of pointer it self, ie value that pointer holds i.e. address where pointer is pointing at.

	// Now play with flags in go

	word := flag.String("name", "aakku", "Enter Your name")
	num := flag.Int("age", 16, "Enter Ur age")
	fmt.Println(word, *word)
	fmt.Println(flag.Parsed())
	flag.Parse()
	fmt.Println(flag.Parsed())
	fmt.Println(word, *word)
	fmt.Println(*num)
	fmt.Println(flag.Args())
}
