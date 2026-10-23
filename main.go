package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CIA-2 Go CI/CD Pipeline Successful! v1.0")
	})
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}