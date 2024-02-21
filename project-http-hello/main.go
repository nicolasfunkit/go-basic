package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls = []string{}
var stats = map[string]int{}

func main() {
	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return // <- this is important!
	}
	calls = append(calls, name)
	stats[name] = stats[name] + 1

	fmt.Fprint(w, "Hello, ", name)
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
}
