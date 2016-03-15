package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir(".")))
        fmt.Println("Serving sample client website on localhost:3000")
        http.ListenAndServe(":3000", nil)
}