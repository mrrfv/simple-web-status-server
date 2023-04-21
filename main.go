package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Immediately close any requests other than GET /check_status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
	})

	// GET /check_status
	http.HandleFunc("/check_status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			defer r.Body.Close()
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
		panic(err)
	}
}
