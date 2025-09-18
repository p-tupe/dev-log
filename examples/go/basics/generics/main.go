// This package shows some of the use cases of generic types.
package main

import (
	"fmt"
	"slices"
)

// To start with, say we want a sum function that only works for either int or float64
func StrictSum[Number int | float64](x, y Number) {
	s := x + y
	fmt.Printf("Sum of %v + %v = %v\n", x, y, s)
}

// But what if we have a custom type based on int?
// It won't work with above Number type
type MyInt int

// For that, we use the "~" operator, which says use the base type
// in this case: any int or any float64
func LessStrictSum[Number ~int | ~float64](x, y Number) {
	s := x + y
	fmt.Printf("Sum of %v + %v = %v\n", x, y, s)
}

// Filter accepts any type of slice.
// Note that if we wrote the function as Filter[T any](arr []T ...
// Then it would accept _any_ type T, including Filter("some string"...
// which is not what we want.
func Filter[Slice ~[]Type, Type any](arr Slice, fn func(x Type) bool) []Type {
	newArr := make([]Type, 0, len(arr))
	for _, x := range arr {
		if fn(x) {
			newArr = append(newArr, x)
		}
	}
	// Clip ensures len(arr) == cap(arr)
	return slices.Clip(newArr)
}

func main() {
	x := 1
	y := 2
	StrictSum(x, y)

	a := 1.4
	b := 2.5
	StrictSum(a, b)

	var m MyInt = 1
	var n MyInt = 2
	// This wont work; MyInt does not satisfy int | float64
	// StrictSum(m, n)
	LessStrictSum(m, n)

	evens := Filter([]int{1, 2, 3, 4, 5}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Even numbers: ", evens)
}
