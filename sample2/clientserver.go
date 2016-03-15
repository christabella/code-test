package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir(".")))
        fmt.Println("Serving sample client website on localhost:3030")
        http.ListenAndServe(":3030", nil)
}