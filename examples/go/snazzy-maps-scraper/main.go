// snazzy-maps-scraper takes in a list of urls of the form https://snazzymaps.com/style/<number>/<name> and downloads the style into <name>.json
//
//	go run . <url1> <url2> ...
package main

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	w := &sync.WaitGroup{}
	urls := os.Args[1:]

	for _, url := range urls {
		w.Add(1)
		go fetchAndParse(url, w)
	}

	w.Wait()
}

func fetchAndParse(url string, w *sync.WaitGroup) {
	defer w.Done()

	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1] + ".json"

	resp, err := http.Get(url)
	panicIfErr(err)

	if resp.StatusCode != 200 {
		panicIfErr(errors.New("Status Code: " + resp.Status))
	}

	outputFile, err := os.Open(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		outputFile, err = os.Create(fileName)
		panicIfErr(err)
	} else {
		panicIfErr(err)
	}
	defer outputFile.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	panicIfErr(err)

	jsonBlock := doc.Find("pre#style-json")
	io.WriteString(outputFile, jsonBlock.Text())
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
