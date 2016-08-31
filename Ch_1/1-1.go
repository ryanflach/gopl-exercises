package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Modify the program to also print os.Args[0]
	fmt.Println(strings.Join(os.Args[0:], " "))
}
