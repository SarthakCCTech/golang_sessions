package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World Again!")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}