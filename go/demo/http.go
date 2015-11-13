// +build OMIT

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", handleHello)
    fmt.Println("serving on http://localhost:8080/hello")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
    log.Println("serving", req.URL)
    fmt.Fprintln(w, "Hello, 世界!")
}