package main

import (
	"fmt"
	"net/http"
)

func handleIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home path")
	})

}

func main() {

	http.HandleFunc("/", handleIndex())
	http.ListenAndServe(":8080", nil)

}
