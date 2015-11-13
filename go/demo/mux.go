// +build OMIT

package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Mux!")
}
