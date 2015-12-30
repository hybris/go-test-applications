package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ok")
}

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", handler)
    if port != "" {
      http.ListenAndServe(":" + port, nil)
    } else {
      http.ListenAndServe(":8080", nil)
    }
}
