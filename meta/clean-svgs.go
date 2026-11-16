package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 || len(os.Args[1:]) == 0 {
		log.Fatalln("No file input list found!")
	}

	// Read a list of paths from std in
	paths := os.Args[1:]

	for _, p := range paths {
		clean(p)
	}
}

var widthRe = regexp.MustCompile(`(?s)(width|height)=".*?"`)
var styleRe = regexp.MustCompile(`(?s)<defs>.*</defs>`)
var strokeRe = regexp.MustCompile(`stroke="#.*?"`)
var fillRe = regexp.MustCompile(`fill="#.*?"`)

func clean(p string) {
	f, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading %s: %s\n", p, err)
	}

	replaced := widthRe.ReplaceAll(f, []byte(""))
	replaced = styleRe.ReplaceAll(replaced, []byte(""))
	replaced = strokeRe.ReplaceAll(replaced, []byte(`stroke="currentColor"`))
	replaced = fillRe.ReplaceAll(replaced, []byte(`fill="currentColor"`))

	err = os.WriteFile(p, replaced, 0644)
	if err != nil {
		log.Fatalf("Error reading %s: %s\n", p, err)
	}
}
