package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func closeTCPConnection(w http.ResponseWriter, r *http.Request) {
	// Get the HTTP hijacker
	hj, ok := w.(http.Hijacker)
	// We can't hijack the connection, so we are closing it normally
	if !ok {
		defer r.Body.Close()
		return
	}
	// Get the underlying TCP connection
	conn, _, err := hj.Hijack()
	// If we can't get the underlying TCP connection, we are closing it normally
	if err != nil {
		defer r.Body.Close()
		return
	}
	// Close the TCP connection
	// Again, if we can't close the connection, we are closing it normally.
	err = conn.Close()
	if err != nil {
		defer r.Body.Close()
	}
}

func main() {
	// Immediately close any requests other than GET /check_status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer closeTCPConnection(w, r)
	})

	// GET /check_status
	http.HandleFunc("/check_status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			defer closeTCPConnection(w, r)
			return
		}

		// Respond with OK if it's a valid request
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Determine port for HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "2384"
	}

	// Start HTTP server
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
