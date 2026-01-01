package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	start := time.Now()
	fmt.Println("Request Received")
	w.Write([]byte("OK"))
	time.Sleep(10 * time.Millisecond)
	duration := time.Since(start) // end timing
	fmt.Printf("%s %s %s took %v\n", r.Method, r.URL.Path, r.UserAgent(), duration)
}

func main() {
	fmt.Println("EdgeProfiler starting...")
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
