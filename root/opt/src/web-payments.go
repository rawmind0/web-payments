package main

import (
    "fmt"
    "net/http"
    "os"
)

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Payments microservice version", os.Getenv("VERSION"))
}

func main() {
	fmt.Println("Starting web-payments version", os.Getenv("VERSION"))
    http.HandleFunc("/", paymentsHandler)
    http.ListenAndServe(":8080", nil)
}
