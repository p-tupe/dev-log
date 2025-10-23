// RoyalRoadScrapper is a cli utility to download full stories
// from royalroad.com.
//
// USAGE:
//
//	go run scrape.go royalroad.com/fiction/x/y/chapter/x/y
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		panic("Error parsing URL!\n\rUSAGE: go run scrape.go <URL>")
	}
	url := os.Args[1]

	outFile, err := os.Create("./story.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for {
		fmt.Println("Fetching: ", url)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		if res.StatusCode != 200 {
			panic("Bad response")
		}
		defer res.Body.Close()
		url = ""

		doc, err := goquery.NewDocumentFromReader(res.Body)
		title := doc.Find("div.row.fic-header > h1").Text()
		content := doc.Find("div.chapter-content").Text()
		outFile.Write([]byte("\n" + title + "\n\n" + content + "\n"))

		for _, s := range doc.Find("a.btn-primary").EachIter() {
			content := s.Text()
			content = strings.ReplaceAll(content, "\n", "\n\n")
			isNextBtn := strings.Contains(content, "Next Chapter")
			if isNextBtn {
				nextUrl, exists := s.Attr("href")
				if exists {
					url = "https://www.royalroad.com" + nextUrl
					break
				}
			}
		}

		if url == "" {
			fmt.Println("Done!")
			break
		}
	}
}
