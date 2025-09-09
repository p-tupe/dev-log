// This program demonstrates the simplest way to use an http-server.
//
// The main route "/" maps to an index.html file,
// while the "/api" routes to a handler function.
// Usage: go run main.go
package main

import (
	"io"
	"net/http"
)

func main() {
	// There is no need to write a new server object,
	// unless we're overriding the defaults.
	// The default http.Handler works just fine for most cases.
	// Usage: curl localhost:8080/
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		// The "/{$}" matches ONLY / route
		// The "/" matches ALL routes (hence used as 404 at the end)
		io.WriteString(w, "Index Page")
	})

	// We can write any data into the write stream.
	// Here's some raw bytes being sent as a response
	// Usage: curl localhost:8080/api
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Raw API Response"))
	})

	// To map a whole file system to a route
	// Note the change from HandleFunc to Handle
	// Usage: curl localhost:8080/static/image.png
	// Usage: curl localhost:8080/static/icon.svg
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Use http.ServeFile(w, r, "./path/to/file") to serve a specific file

	// And to use a handler for more granular control
	// Usage: curl localhost:8080/handler
	http.Handle("/handler", NewHndlr())

	// This is the "catch all" route to send a 404
	// Usage: curl localhost:8080/unknown
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "404 Not Found")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// A handler must have a ServeHTTP method,
// such that is satisfies the http.Handler interface.
func NewHndlr() *hndlr {
	return &hndlr{}
}

type hndlr struct {
	// add dependencies here
}

func (h *hndlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("From inside handler"))
}
