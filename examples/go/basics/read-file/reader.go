// This program demonstrates different ways to read a file
//
// Usage: `go run main.go`
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readWhole()
	readLines()
}

// Reads a whole file into memory
func readWhole() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Whole file: ", string(f))
}

// Reads a file one line at a time
func readLines() {
	f, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}

	b := bufio.NewScanner(f)
	for b.Scan() {
		l := b.Text()
		fmt.Println("Line: ", l)
	}
	if err := b.Err(); err != nil {
		panic(err)
	}
}
