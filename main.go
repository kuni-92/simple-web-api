package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", apiHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error. process exit.")
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "Hello!!! You have visited %d times !", count)
}
