package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service 1 (Go) OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Go service running on port 9090...")
	http.ListenAndServe(":9090", nil)
}
