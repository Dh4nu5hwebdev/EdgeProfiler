package main

import (
	"fmt"
	"net/http"
	"time"
)

func backendHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(50 * time.Millisecond)
	fmt.Println("Request intercepted.")
	w.Write([]byte("Intercepted"))

}

func main() {
	fmt.Println("Backend Server Program executing..")
	http.HandleFunc("/", backendHandler)
	fmt.Println("Server running on port: 9000")
	http.ListenAndServe(":9000", nil)
}
