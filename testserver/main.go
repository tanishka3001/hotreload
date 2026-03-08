package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server UPDATED....")
	})

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}