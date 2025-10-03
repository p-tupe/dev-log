// This package contains some string manipulation functions
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(addExtraNewLines("Line 1\nLine 2"))
}

// This function shows the "replaceAll" util that
// takes in a string and adds an extra newline for evey present newline.
//
// If ip = "Line 1\nLine 2" then op = "Line 1\n\nLine 2"
func addExtraNewLines(ip string) (op string) {
	return strings.ReplaceAll(ip, "\n", "\n\n")
}
