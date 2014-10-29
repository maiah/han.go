package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}
