// This is a simple mail server implementation that accepts an email body content
// and sends it over via SMTP using go's builtin [net/smtp](https://pkg.go.dev/net/smtp) package.
package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// Must update before running
var (
	Username = ""
	Password = ""
	Host     = "smtp.mail.com"
	Port     = "587"
	From     = "from@mail.com"
	To       = []string{"to@mail.com"}

	// It is recommended to generate this random string before deploying the app
	// Can use: `head -n1 /dev/urandom | base64`
	AuthKey = "some-random-string"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next(w, r)
	}
}

func main() {
	if Username == "" || Password == "" || Host == "" || Port == "" || From == "" || To[0] == "" {
		slog.Error("Required variables not found!")
		os.Exit(1)
	}

	auth := smtp.PlainAuth("", Username, Password, Host)

	http.HandleFunc("/", middleware(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	http.HandleFunc("POST /", middleware(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Authorization")

		if key != AuthKey {
			slog.Error("Received invalid Authorization key: " + key)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Unable to read request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := []byte("From: " + From + "\r\nTo: " + strings.Join(To, ",") + "\r\nSubject: Automated Email\r\n\r\n" + string(body) + "\r\n")

		slog.Info("Sending: " + string(body))

		err = smtp.SendMail(Host+":"+Port, auth, From, To, msg)
		if err != nil {
			slog.Error("Unable to send email: " + err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}))

	slog.Info("Mail Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
