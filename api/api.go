package api

import (
	"fmt"
	"net/http"
)

var count int

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[HandlerMain] Start.")
	count++
	fmt.Fprintf(w, "Hello!!! You have visited %d times !", count)
	fmt.Println("[HandlerMain] Finish.")
}
