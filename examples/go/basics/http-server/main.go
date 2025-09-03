// This program demonstrates the simplest way to use an http-server.
//
// The main route "/" maps to an index.html file,
// while the "/api" routes to a handler function.
// Usage: go run main.go
package main

import (
	"net/http"
)

func main() {
	// There is no need to write a new server object,
	// unless we're overriding the defaults.
	// The default http.Handler works just fine for most cases.

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

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
