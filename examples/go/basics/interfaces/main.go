// This package describles how interfaces and concrete types
// are tied to each other
//
//	Usage: go run main.go
package main

import (
	"fmt"
)

// Interface is just a set of behaviours/traits/methods/function signatures.
type Calc interface {
	Add() int
	Sub() int
}

// Any type that implements those methods satisfies the interface.
type X int

// Like here, x satisfies Calc interface by implementing
// Add and Sub methods.
func (x X) Add() int {
	return int(x) + 1
}

func (x X) Sub() int {
	return int(x) - 1
}

// Here's another type that satisfies X
type S struct {
	value int
}

func (s S) Add() int {
	return s.value + 1
}

func (s S) Sub() int {
	return s.value - 1
}

// We usually write functions that accept an interface
func MyCalc(c Calc) int {
	return c.Add() + c.Sub()
}

// Functions usually return a concrete type
func AnotherCalc(c Calc) S {
	return S{value: MyCalc(c)}
}

// Most interfaces have a single method
type Hndler interface {
	Handle(input string) string
}

// This will substitue as the Handle func
type HndlerFunc func(string) string

// Where we just write a thin wrapper on top
func (f HndlerFunc) Handle(s string) string {
	return f(s)
}

// And thus, we can write functions that satisfy the interface
// without even writing methods specifically for it
func ConsumerHandlerFunc(h HndlerFunc) {
	h.Handle("xyz")
}

func main() {
	// We can check if a type satisfies an interface using a type check
	var y X = 5
	_, ok := any(y).(Calc)
	fmt.Println("y satisfies X: ", ok) // true
	fmt.Println(MyCalc(y))

	var z int = 9
	_, ok = any(z).(Calc)
	fmt.Println("z satisfies X: ", ok) // false

	// For the functions that expect an interface,
	// We can pass any type that satisfies it.
	ss := S{value: 3}
	fmt.Println(ss.Add())
	fmt.Println(ss.Sub())
	fmt.Println(MyCalc(ss))

	// This will satisfy the consumer function that expects
	// a Hndler, without having to implement Handle explicitly
	ConsumerHandlerFunc(func(s string) string {
		return "whatever"
	})
}
