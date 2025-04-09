package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    log.Println("Starting Lambda Web Adapter mock server on :8080")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Header().Set("Connection", "keep-alive")
        w.Header().Set("Keep-Alive", "timeout=72")
        w.WriteHeader(200)
        fmt.Fprintln(w, "Adapter mock response with keep-alive headers")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
