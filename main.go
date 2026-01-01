package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	start := time.Now()
	backendReq, err := http.NewRequest(
		r.Method,
		"http://localhost:9000"+r.URL.Path,
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to create backend request", 500)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(backendReq)
	if err != nil {
		http.Error(w, "Backend unreachable", 502)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start) // end timing
	w.WriteHeader(resp.StatusCode)
	fmt.Fprintln(w, "Response from backend")

	fmt.Printf(
		"%s %s -> backend took %v\n",
		r.Method,
		r.URL.Path,
		duration,
	)
}

func main() {
	fmt.Println("EdgeProfiler starting...")
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
