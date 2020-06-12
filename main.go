package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello World")
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
	}
	fmt.Println("Test")
	fmt.Println(s)
}
