package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    //Prints guessing game with the h1 tags into the console
    fmt.Fprintf(w, "<h1>Guessing game!</h1>")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}