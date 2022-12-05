package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	calls []string
	stats = make(map[string]int)
)

func main() {
	http.HandleFunc("/hello", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s", name)

	calls = append(calls, r.URL.Query().Get("name"))
	stats[r.URL.Query().Get("name")]++

	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n", stats)
}
