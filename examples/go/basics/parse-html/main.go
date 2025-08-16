// This program demonstrates a convenient method to
// parse html files using an external package.
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.ReadFile("tasks.html")
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(file)
	doc, err := goquery.NewDocumentFromReader(buffer)

	if err != nil {
		panic(err)
	}

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
