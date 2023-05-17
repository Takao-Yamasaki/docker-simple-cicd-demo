package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		fmt.Fprintf(w, "Hello World!!!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
