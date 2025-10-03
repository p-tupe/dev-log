// This package starts a server that receives a POST request with a given AUTH_TOKEN
// and generates a slack message using slack WEBHOOK_URL.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NOTE: These are dummy values
const (
	WEBHOOK_URL = "https://hooks.slack.com/services/xxx/yyy/zzz"
	AUTH_TOKEN  = "randomly-generated-string"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.Handle("/", corsMiddleware(http.HandlerFunc(handler)))
	fmt.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != AUTH_TOKEN {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	var msg Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received message to send %s\n", msg.Text)

	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	resp, err := http.Post(WEBHOOK_URL, "application/json", bytes.NewReader(jsonData))
	io.WriteString(w, resp.Status)
}
