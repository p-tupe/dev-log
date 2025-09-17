// This example demonstrates the idiomatic way of generating
// enums in go.
//
// Before running `go run .` do a quick `go generate` to create
// required files using stringer.
package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Season
type Season int

const (
	// iota is a keyword that starts at x (default 0)
	// and increments at every step
	Unknown Season = iota

	// So winter will be 1
	Winter
	// Summer will be 2, and so on
	Summer
	Monsoon
	Fall
	Spring
)

func main() {
	s := Summer
	fmt.Println(s)

	var x Season
	fmt.Printf("%v %T", x, x)
}
