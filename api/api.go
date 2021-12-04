package api

import (
	"encoding/json"
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

type Response struct {
	Key   string
	Value string
}

func HandlerJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[HandlerJson] Start.")

	query := r.URL.Query()
	k := query.Get("key")
	v := query.Get("value")

	res, err := json.Marshal(Response{Key: k, Value: v})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad query parameter."))
		fmt.Println("[HandlerJson] Error.")
		return
	}

	fmt.Fprint(w, string(res))
	fmt.Println("[HandlerJson] Finish.")
}
