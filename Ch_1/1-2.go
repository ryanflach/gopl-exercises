package main

import (
	"fmt"
	"os"
)

func main() {
	// Modify the program to print the index and value of each arg - 1 per line
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
