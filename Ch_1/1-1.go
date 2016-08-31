package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Modify the program to also print os.Args[0]
	fmt.Println(strings.Join(os.Args[0:], " "))
	// Modify the program to print the index and value of each arg - 1 per line
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
