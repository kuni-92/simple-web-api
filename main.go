package main

import (
	"fmt"
	"net/http"

	"github.com/kunikuni03/simple-web-api/api"
)

func main() {
	http.HandleFunc("/api", api.HandlerMain)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error. process exit.")
	}
}
